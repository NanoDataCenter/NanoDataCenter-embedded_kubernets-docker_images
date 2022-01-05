package eto_monitor

import (
	"fmt"
    
	"lacima.com/cf_control"
	"lacima.com/go_application_containers/irrigation/eto/eto_calc/eto_support"
	"lacima.com/redis_support/generate_handlers"
	//"lacima.com/redis_support/graph_query"
	"lacima.com/redis_support/redis_handlers"
	"time"
    "lacima.com/Patterns/msgpack_2"
    "lacima.com/Patterns/logging_support"
    "encoding/json"
)

type ETO_Data_Lookup_Type func(string)[]eto_support.ETO_INPUT





var CF_site_node_control_cluster cf.CF_CLUSTER_TYPE
var Eto_data_structs *map[string]interface{}

/*
 *
 *  Redis data structures
 *
 */
var eto_control                   redis_handlers.Redis_Hash_Struct
var eto_rollover_log              *logging_support.Incident_Log_Type
var eto_update_log               *logging_support.Incident_Log_Type




func Start() {

	
	CF_site_node_control_cluster.Cf_cluster_init()
	CF_site_node_control_cluster.Cf_set_current_row("eto_monitor")

	
	setup_data_structures()
    
	setup_chain_flow(&CF_site_node_control_cluster)
	
	go execute()

}

func execute() {
	(CF_site_node_control_cluster).CF_Fork()
}





func setup_data_structures() {
    eto_rollover_log    = logging_support.Construct_incident_log([]string{"WEATHER_DATA_STRUCTURE:WEATHER_DATA_STRUCTURE","ROLL_OVER:ROLL_OVER" ,"INCIDENT_LOG"} ) 
    eto_update_log      = logging_support.Construct_incident_log([]string{"WEATHER_DATA_STRUCTURE:WEATHER_DATA_STRUCTURE","ETO_UPDATE:ETO_UPDATE" ,"INCIDENT_LOG"} ) 
	search_list := []string{"WEATHER_DATA"}
	Eto_data_structs = data_handler.Construct_Data_Structures(&search_list)
	eto_control      = (*Eto_data_structs)["ETO_CONTROL"].(redis_handlers.Redis_Hash_Struct)
	
}




        
        
func check_new_day_rollover(system interface{}, chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {
    
    
    message_pack_false := msg_pack_utils.Pack_bool(false)
    error_flag         := true
    
    if eto_control.HGet("ETO_UPDATE_FLAG") != message_pack_false {
          error_flag = false
    }
        
    if eto_control.HGet("ETO_LOG_FLAG") != message_pack_false {
          error_flag = false
    }
    fmt.Println("**************************** eto bin update",error_flag)
   if error_flag == false {
        state := make(map[string]interface{})
        state["system"]              = "eto_monitor"
        state["subsystem"]           = "new_day_rollover"
        post_incident_day_rollover(state)
    }
    return cf.CF_DISABLE
}
             

func check_eto_bin_update(system interface{}, chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {
    
    
    message_pack_true := msg_pack_utils.Pack_bool(true)
    error_flag         := true
    if eto_control.HGet("ETO_UPDATE_FLAG") != message_pack_true {
          error_flag = false
    }
        
    if eto_control.HGet("ETO_LOG_FLAG") != message_pack_true {
          error_flag = false
    }
    fmt.Println("**************************** eto update state",error_flag)
    
    
    if error_flag == false {
        state := make(map[string]interface{})
        state["system"]              = "eto_monitor"
        state["subsystem"]           = "eto_update"
        post_incident_eto_update(state)
    }
    return cf.CF_DISABLE
}
 


func setup_chain_flow(cf_cluster *cf.CF_CLUSTER_TYPE) {

	var cf_control cf.CF_SYSTEM_TYPE

	cf_control.Init(cf_cluster, "ETO_MONITOR", true, time.Second)

	
    
    cf_control.Add_Chain("Monitor_day_rollover", true)
    cf_control.Cf_add_log_link("monitor rollover")
    cf_control.Cf_wait_hour_minute_le(7, 15)
    cf_control.Cf_wait_hour_minute_ge(7, 45)
    cf_control.Cf_add_one_step(check_new_day_rollover, make(map[string]interface{}))
    cf_control.Cf_wait_hour_minute_ge(8, 0)
    cf_control.Cf_add_reset()

    cf_control.Add_Chain("monitor_eto_update", true)
    cf_control.Cf_add_log_link("monitor update")
    cf_control.Cf_wait_hour_minute_le(13, 0)
    cf_control.Cf_wait_hour_minute_ge(13, 1)
    cf_control.Cf_add_one_step(check_eto_bin_update, make(map[string]interface{})) 
    cf_control.Cf_wait_hour_minute_ge(14, 0)
    cf_control.Cf_add_reset()



}


func post_incident_day_rollover(incident_data map[string]interface{}){
    
    request_json,err := json.Marshal(&incident_data)
    if err != nil{
          panic("json marshall error")
    }  
    fmt.Println("request_json",string(request_json))
    eto_rollover_log.Log_data(string(request_json))
}

func post_incident_eto_update(incident_data map[string]interface{}){
    
    request_json,err := json.Marshal(&incident_data)
    if err != nil{
          panic("json marshall error")
    }  
    fmt.Println("request_json",string(request_json))
    eto_update_log.Log_data(string(request_json))
}
