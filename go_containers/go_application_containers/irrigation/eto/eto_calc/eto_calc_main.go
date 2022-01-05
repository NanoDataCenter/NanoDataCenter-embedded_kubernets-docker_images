package eto_calc

import (
	//"fmt"
    
	"lacima.com/cf_control"
	"lacima.com/go_application_containers/irrigation/eto/eto_calc/eto_support"
	"lacima.com/redis_support/generate_handlers"
	//"lacima.com/redis_support/graph_query"
	"lacima.com/redis_support/redis_handlers"
	"time"
    "lacima.com/Patterns/msgpack_2"
    "lacima.com/server_libraries/postgres"
)

type ETO_Data_Lookup_Type func(string)[]eto_support.ETO_INPUT





var CF_site_node_control_cluster cf.CF_CLUSTER_TYPE
var Eto_data_structs *map[string]interface{}

/*
 *
 *  Redis data structures
 *
 */
var eto_control           redis_handlers.Redis_Hash_Struct
var eto_exceptions        redis_handlers.Redis_Hash_Struct

var eto_data              redis_handlers.Redis_Hash_Struct
var rain_data             redis_handlers.Redis_Hash_Struct
var eto_stream_data       redis_handlers.Redis_Hash_Struct
var eto_history           pg_drv.Postgres_Stream_Driver
var rain_history          pg_drv.Postgres_Stream_Driver
/*
 * 
 *  Postgres Data Structures
 * 
 * 
 */

    





func Start(site_data map[string]interface{}) {

	eto_support.Init_Secret_Support(site_data)
	CF_site_node_control_cluster.Cf_cluster_init()
	CF_site_node_control_cluster.Cf_set_current_row("eto_calc")

	
	setup_data_structures()
    eto_support.Init_Record_Store(eto_data, rain_data,eto_stream_data)
    initialize_eto_calc_setup()
	setup_eto_calculators()
	setup_chain_flow(&CF_site_node_control_cluster)
	
	go execute()

}

func execute() {
	(CF_site_node_control_cluster).CF_Fork()
}





func setup_data_structures() {

	search_list := []string{"WEATHER_DATA"}
	Eto_data_structs = data_handler.Construct_Data_Structures(&search_list)
	eto_control      = (*Eto_data_structs)["ETO_CONTROL"].(redis_handlers.Redis_Hash_Struct)
	eto_exceptions   = (*Eto_data_structs)["EXCEPTION_VALUES"].(redis_handlers.Redis_Hash_Struct)
	
	eto_data         = (*Eto_data_structs)["ETO_VALUES"].(redis_handlers.Redis_Hash_Struct)
	rain_data        = (*Eto_data_structs)["RAIN_VALUES"].(redis_handlers.Redis_Hash_Struct)
    eto_stream_data  = (*Eto_data_structs)["ETO_STREAM_DATA"].(redis_handlers.Redis_Hash_Struct)
    eto_history      = (*Eto_data_structs)["ETO_HISTORY"].(pg_drv.Postgres_Stream_Driver)
    rain_history     = (*Eto_data_structs)["RAIN_HISTORY"].(pg_drv.Postgres_Stream_Driver)
}



func setup_chain_flow(cf_cluster *cf.CF_CLUSTER_TYPE) {

	var cf_control cf.CF_SYSTEM_TYPE

	cf_control.Init(cf_cluster, "ETO_CALCULATOR", false, time.Second)

	//**********************************************************************************
	cf_control.Add_Chain("test_generator", true)
	cf_control.Cf_add_one_step(new_day_rollover, make(map[string]interface{}))
    cf_control.Cf_add_enable_chains_links([]string{"eto_make_measurements"})
    cf_control.Cf_add_wait_interval(time.Minute * 5)
	cf_control.Cf_add_enable_chains_links([]string{"update_eto_bins"})
	cf_control.Cf_add_terminate()
    
    
    cf_control.Add_Chain("new_day_rollover", true)
	cf_control.Cf_add_log_link("start new day rollover ")
	cf_control.Cf_wait_hour_minute_le(7, 0) // Wait till time is less than 7
	cf_control.Cf_wait_hour_minute_ge(7, 1) // Then when time gt 7 act
	cf_control.Cf_add_one_step(new_day_rollover, make(map[string]interface{}))
	cf_control.Cf_add_reset()
	//----------------------------------------------------------------------------------
	//***********************************************************************************
	cf_control.Add_Chain("enable_measurement", true)
	cf_control.Cf_add_log_link("starting enable_measurement")
	
    
	cf_control.Cf_wait_hour_minute_le(8, 0) // Wait till time is less than 8
	cf_control.Cf_wait_hour_minute_ge(8, 0) // Then when time turns 8 act
	cf_control.Cf_add_enable_chains_links([]string{"eto_make_measurements"})
	cf_control.Cf_add_log_link("enabling making_measurement")
	cf_control.Cf_wait_hour_minute_ge(11, 0)
	cf_control.Cf_add_enable_chains_links([]string{"update_eto_bins"})
	cf_control.Cf_add_log_link("enable_update_eto_bins")
	cf_control.Cf_wait_hour_minute_ge(23, 0)
    cf_control.Cf_add_log_link("terminating measurement")
	cf_control.Cf_add_disable_chains_links([]string{"eto_make_measurements", "update_eto_bins"})
	cf_control.Cf_wait_hour_minute_le(1, 0)
	cf_control.Cf_add_reset()
	//----------------------------------------------------------------------------------

	//***********************************************************************************
	cf_control.Add_Chain("update_eto_bins", false)
    cf_control.Cf_add_log_link("starting update eto bins")
	cf_control.Cf_add_wait_interval(time.Minute * 8)
	cf_control.Cf_add_log_link("updating eto bins")
	cf_control.Cf_add_unfiltered_element(update_eto_bins, make(map[string]interface{}))
	cf_control.Cf_add_disable_chains_links([]string{"eto_make_measurements"})
    cf_control.Cf_add_log_link("terminating eto_bins")
	cf_control.Cf_add_terminate()

	//----------------------------------------------------------------------------------

	//***********************************************************************************
	cf_control.Add_Chain("eto_make_measurements", false)
	cf_control.Cf_add_log_link("starting make measurement")
	cf_control.Cf_add_one_step(make_measurement, make(map[string]interface{}))
	cf_control.Cf_add_wait_interval(time.Minute * 5)
	cf_control.Cf_add_log_link("Receiving minute tick")
	cf_control.Cf_add_reset()

}
func new_day_rollover(system interface{}, chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

    message_pack_false := msg_pack_utils.Pack_bool(false)
    message_pack_zero  := msg_pack_utils.Pack_float64(0)
    eto_control.Delete_All()
    eto_data.Delete_All()
    rain_data.Delete_All()             
    eto_stream_data.Delete_All()	
	
	eto_control.HSet("ETO_UPDATE_FLAG",message_pack_false)
	eto_control.HSet("ETO_LOG_FLAG",message_pack_false)
	eto_control.HSet("ETO_UPDATE_VALUE",message_pack_zero)
	return cf.CF_DISABLE
}

func make_measurement(system interface{}, chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {
    do_calculators()
	return cf.CF_DISABLE
}


func update_eto_bins(system interface{}, chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {
      message_pack_true := msg_pack_utils.Pack_bool(true)
      if eto_control.HGet("ETO_UPDATE_FLAG") == message_pack_true {
          return cf.CF_DISABLE
      }
      eto,status := find_eto_lowest_priority()
      if status == false {
          return cf.CF_HALT
      }
      
      eto_control.HSet("ETO_UPDATE_FLAG",message_pack_true)
      eto_msgpack  := msg_pack_utils.Pack_float64(eto)
      eto_control.HSet("ETO_UPDATE_VALUE",eto_msgpack)
      //eto_support.Update_Accumulation_Tables(eto_accumulation, eto)
      log_eto_data()
      log_rain_data()
      return cf.CF_DISABLE
}


func log_eto_data(){
    data := eto_data.HGetAll()
    
    for key,value := range data {
        eto_history.Insert( key,"","","","",value )
        
    }

}
func log_rain_data(){
    data := rain_data.HGetAll()
    
    for key,value := range data {
        rain_history.Insert( key,"","","","",value )
        
    }

}

func find_eto_lowest_priority()(float64,bool){
    var output eto_support.ETO_RAIN_TYPE
    status := false
    output.Priority = 1E9
    output.Value    = 0
    
    data := eto_data.HKeys()
    for _,key := range data{
        status = true
        eto_data := eto_support.ETO_HGet(key)
        if eto_data.Priority < output.Priority {
            output.Priority = eto_data.Priority
            output.Value    = eto_data.Value
        }
        
    }
    return output.Value,status
    
}


