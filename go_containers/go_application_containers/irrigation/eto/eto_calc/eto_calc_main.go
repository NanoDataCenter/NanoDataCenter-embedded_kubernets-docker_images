package eto_calc

import (
    "fmt"
    "lacima.com/cf_control"
    "lacima.com/go_application_containers/irrigation/eto/eto_calc/eto_support"
    "time"
     "lacima.com/redis_support/graph_query"
    "lacima.com/redis_support/generate_handlers"
    "lacima.com/redis_support/redis_handlers"
    
)
    

type ETO_CALC_SETUP_TYPE func(string,string,map[string]string,redis_handlers.Redis_Hash_Struct,redis_handlers.Redis_Hash_Struct)interface{}

type ETO_CALC_TYPE func(interface{})

var  CF_site_node_control_cluster cf.CF_CLUSTER_TYPE
var  Eto_data_structs  *map[string]interface{}
/*
 * 
 *  Redis data structures
 * 
 */
var  eto_control       redis_handlers.Redis_Hash_Struct
var  eto_exceptions    redis_handlers.Redis_Hash_Struct
var  eto_accumulation  redis_handlers.Redis_Hash_Struct
var  eto_data          redis_handlers.Redis_Hash_Struct
var  rain_data         redis_handlers.Redis_Hash_Struct
/*
 * Postgres database handlers
 * 
 * 
 */
//su.Cd_Rec.Create_postgres_stream( "ETO_HISTORY","admin","password","admin",2*365*24*3600)
//su.Cd_Rec.Create_postgres_stream( "RAIN_HISTORY","admin","password","admin",2*365*24*3600)
//su.Cd_Rec.Create_postgres_stream( "EXCEPTION_LOG","admin","password","admin",2*365*24*3600)





var eto_calc_setup map[string]ETO_CALC_SETUP_TYPE

var eto_calc_compute map[string]ETO_CALC_TYPE

var eto_calculators map[string]interface{}



func Start(site_data map[string]interface{}){
    
    eto_support.Init_Secret_Support(site_data)
    CF_site_node_control_cluster.Cf_cluster_init()
	CF_site_node_control_cluster.Cf_set_current_row("site_node_control")
    
    initialize_eto_calc_setup()
    setup_data_structures()
    
    
    eto_support.Setup_ETO_Accumulation_Table(eto_accumulation)
   
    setup_eto_calculators()
    //setup_chain_flow(&CF_site_node_control_cluster)
    panic("done")
    go execute()   
    
}

func execute(){
    (CF_site_node_control_cluster).CF_Fork()
}

func initialize_eto_calc_setup(){
    eto_calculators = make(map[string]interface{})
    eto_calc_setup = make(map[string]ETO_CALC_SETUP_TYPE)
    eto_calc_setup["MESSO_ETO"]              = messo_eto_setup 
    eto_calc_setup["CIMIS"]                  = cimis_setup
    eto_calc_setup["MESSO_RAIN"]             = messo_rain_setup
    eto_calc_setup["CIMIS_SAT"]              = cimis_sat_setup
    eto_calc_setup["WUNDERGROUND"]           = wunderground_setup
    eto_calc_setup["HYBRID_SETUP"]           = hybrid_setup
    eto_calc_compute = make(map[string]ETO_CALC_TYPE)
    eto_calc_compute["MESSO_ETO"]              = messo_eto_calc
    eto_calc_compute["CIMIS"]                  = cimis_calc
    eto_calc_compute["MESSO_RAIN"]             = messo_rain_calc
    eto_calc_compute["CIMIS_SAT"]              = cimis_sat_calc
    eto_calc_compute["WUNDERGROUND"]           = wunderground_calc
    eto_calc_compute["HYBRID_SETUP"]           = hybrid_calc
    
}
/*
messo_rain_setup map[access_key:"MESSOWEST" name:"SRUC1_RAIN" namespace:"[SYSTEM:farm_system][SITE:LACIMA_SITE][WEATHER_STATIONS:WEATHER_STATIONS][WEATHER_STATION:SRUC1_RAIN]" priority:3 station:"SRUC1" sub_id:"" type:"MESSO_RAIN" url:"http://et.water.ca.gov/api/data"]
messo_eto_setup map[access_key:"MESSOWEST" altitude:2400 latitude:33.578156 name:"SRUC1" namespace:"[SYSTEM:farm_system][SITE:LACIMA_SITE][WEATHER_STATIONS:WEATHER_STATIONS][WEATHER_STATION:SRUC1]" priority:3 station:"SRUC1" sub_id:"" type:"MESSO_ETO" url:"http://et.water.ca.gov/api/data"]
cimis_sat_setup map[access_key:"ETO_CIMIS_SATELLITE" latitude:33.578156 longitude:-117.29945 name:"ETO_CIMIS_SATELLITE" namespace:"[SYSTEM:farm_system][SITE:LACIMA_SITE][WEATHER_STATIONS:WEATHER_STATIONS][WEATHER_STATION:ETO_CIMIS_SATELLITE]" priority:4 sub_id:"" type:"CIMIS_SAT" url:"http://et.water.ca.gov/api/data"]
wunderground_setup map[access_key:"WUNDERGROUND" alt:2400 lat:33.2 name:"WUNDERGROUND" namespace:"[SYSTEM:farm_system][SITE:LACIMA_SITE][WEATHER_STATIONS:WEATHER_STATIONS][WEATHER_STATION:WUNDERGROUND]" priority:1000 sub_id:"KCAMURRI101" type:"WUNDERGROUND"]
cimis_setup map[access_key:"ETO_CIMIS" altitude:2400 name:"ETO_CIMIS" namespace:"[SYSTEM:farm_system][SITE:LACIMA_SITE][WEATHER_STATIONS:WEATHER_STATIONS][WEATHER_STATION:ETO_CIMIS]" priority:2 station:"62" sub_id:"" type:"CIMIS" url:"http://et.water.ca.gov/api/data"]
hybrid_setup map[main:"WUNDERGROUND" name:"WUNDERGROUND_HYBRID" namespace:"[SYSTEM:farm_system][SITE:LACIMA_SITE][WEATHER_STATIONS:WEATHER_STATIONS][WEATHER_STATION:WUNDERGROUND_HYBRID]" priority:4 solar:"SRUC1" sub_id:"KCAMURRI101" type:"HYBRID_SETUP"]
panic: done
*/

func messo_eto_setup(key, calc_type string,input map[string]string,eto_data,rain_data redis_handlers.Redis_Hash_Struct)interface{}{
    //fmt.Println("messo_eto_setup",input)
    return_value := make(map[string]interface{}) 
    
    return return_value
}

func cimis_setup(key, calc_type string,input map[string]string,eto_data,rain_data redis_handlers.Redis_Hash_Struct)interface{}{
    fmt.Println("cimis_setup",input)

    return eto_support.Create_CIMIS_ETO(key, calc_type, input ,eto_data,rain_data )
}
func messo_rain_setup(key, calc_type string,input map[string]string,eto_data,rain_data redis_handlers.Redis_Hash_Struct)interface{}{
    //fmt.Println("messo_rain_setup",input)
    return_value := make(map[string]interface{}) 
    return return_value
}
func cimis_sat_setup(key, calc_type string, input map[string]string,eto_data,rain_data redis_handlers.Redis_Hash_Struct)interface{}{
    //fmt.Println("cimis_sat_setup",input)
   return_value := make(map[string]interface{}) 
    return return_value
}

func wunderground_setup(key, calc_type string,input map[string]string,eto_data,rain_data redis_handlers.Redis_Hash_Struct)interface{}{
    //fmt.Println("wunderground_setup",input)
   return_value := make(map[string]interface{}) 
    return return_value
}
func hybrid_setup(key , calc_type string,input map[string]string,eto_data,rain_data redis_handlers.Redis_Hash_Struct)interface{}{
    //fmt.Println("hybrid_setup",input)
   return_value := make(map[string]interface{}) 
    return return_value
}

func messo_eto_calc(input interface{}){
    fmt.Println("messo_eto_calc")
    
}

func cimis_calc(input interface{}){
    fmt.Println("cimis_calc")
    data := input.(eto_support.CIMIS_ETO_TYPE)
    data.Compute_eto()
}
func messo_rain_calc(input interface{}){
    fmt.Println("messo_rain_calc")
}
func cimis_sat_calc(input interface{}){
    fmt.Println("cimis_sat_calc")
   
}
func wunderground_calc(input interface{}){
    fmt.Println("wunderground_calc")
   
}
func hybrid_calc(input interface{}){
    fmt.Println("hybrid_calc")
   
}





func setup_data_structures(){
                         
    search_list := []string{"WEATHER_DATA"}
    Eto_data_structs   =  data_handler.Construct_Data_Structures(&search_list)
    eto_control        = (*Eto_data_structs)["ETO_CONTROL"].(redis_handlers.Redis_Hash_Struct)
    eto_exceptions     = (*Eto_data_structs)["EXCEPTION_VALUES"].(redis_handlers.Redis_Hash_Struct)
    eto_accumulation   = (*Eto_data_structs)["ETO_ACCUMULATION_TABLE"].(redis_handlers.Redis_Hash_Struct)
    eto_data           = (*Eto_data_structs)["ETO_VALUES"].(redis_handlers.Redis_Hash_Struct)
    rain_data          = (*Eto_data_structs)["RAIN_VALUES"].(redis_handlers.Redis_Hash_Struct)
    
}



func setup_eto_calculators(){
    search_list := []string{"WEATHER_STATIONS:WEATHER_STATIONS" , "WEATHER_STATION"}
    nodes := graph_query.Common_qs_search(&search_list)
    for _,node := range nodes{
        ws_type := graph_query.Convert_json_string(node["type"])
        sub_id  := graph_query.Convert_json_string(node["sub_id"])
        value,ok := eto_calc_setup[ws_type]
        if ok == false {
            panic("bad calculator type")
        }
        key := ws_type+":"+sub_id
        fmt.Println("key",key)
        eto_calculators[ws_type]=value(key,ws_type,node,eto_data,rain_data)
        
    }
    fmt.Println("eto_calculators",eto_calculators)
    
    do_calculators()
    panic("done")
}

func do_calculators(){
    for ws_type,data := range eto_calculators {
        
        eto_calc_compute[ws_type](data)

        
    }
}



func setup_chain_flow(cf_cluster *cf.CF_CLUSTER_TYPE){
    
    var cf_control  cf.CF_SYSTEM_TYPE

    cf_control.Init(cf_cluster ,"ETO_CALCULATOR",true, time.Second)
    
    
     //**********************************************************************************
     cf_control.Add_Chain("test_generator",false)
     cf_control.Cf_add_enable_chains_links([]string{"eto_make_measurements"})
     cf_control.Cf_add_one_step(new_day_rollover,make(map[string]interface{}))
     cf_control.Cf_add_enable_chains_links([]string{"update_eto_bins"})
     cf_control.Cf_add_terminate()
     
     
     //----------------------------------------------------------------------------------
     //***********************************************************************************
     cf_control.Add_Chain("enable_measurement",true)
     cf_control.Cf_add_log_link("starting enable_measurement")
     cf_control.Cf_wait_hour_minute_le(7,0)  // Wait till time is less than 8
     cf_control.Cf_wait_hour_minute_ge( 7,0 ) // Then when time turns 8 act    
     cf_control.Cf_add_one_step(new_day_rollover,make(map[string]interface{}))
     cf_control.Cf_wait_hour_minute_le(8,0)  // Wait till time is less than 8
     cf_control.Cf_wait_hour_minute_ge (  8,0 ) // Then when time turns 8 act
     cf_control.Cf_add_enable_chains_links([]string{"eto_make_measurements"})
     cf_control.Cf_add_log_link("enabling making_measurement")
     cf_control.Cf_wait_hour_minute_ge(11,0)
     cf_control.Cf_add_enable_chains_links([]string{"update_eto_bins"})
     cf_control.Cf_add_log_link("enable_update_eto_bins")
     cf_control.Cf_wait_hour_minute_ge(   23,0 )
     cf_control.Cf_add_disable_chains_links([]string{"eto_make_measurements","update_eto_bins"})
     cf_control.Cf_wait_hour_minute_le(  1,0 )
     cf_control.Cf_add_reset( )
     //----------------------------------------------------------------------------------
      
 
   
     //***********************************************************************************
     cf_control.Add_Chain("update_eto_bins",false)
     cf_control.Cf_add_wait_interval( time.Minute*8)
     cf_control.Cf_add_log_link("updating eto bins")
     cf_control.Cf_add_unfiltered_element(update_eto_bins,make(map[string]interface{}) )
     cf_control.Cf_add_disable_chains_links([]string{"eto_make_measurements"})
     cf_control.Cf_add_terminate()
   
     //----------------------------------------------------------------------------------
    
     //***********************************************************************************
     cf_control.Add_Chain("eto_make_measurements",false)
     cf_control.Cf_add_log_link("starting make measurement")
     cf_control.Cf_add_one_step(make_measuremeent,make(map[string]interface{}))
     cf_control.Cf_add_wait_interval(time.Minute*5)
     cf_control.Cf_add_log_link("Receiving minute tick")
     cf_control.Cf_add_reset( )
    
   
   
   
 
   
}  
func new_day_rollover( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{
  
    //ds_handlers["EXCEPTION_VALUES"].Delete_Sll()
    //ds_handlers["ETO_VALUES"].Delete_Sll()
    //ds_handlers["RAIN_VALUES"].Delete_All()
    //ds_handlers["ETO_CONTROL"].HSET("ETO_UPDATE_FLAG",0)
    //ds_handlers["ETO_CONTROL"].HSET("ETO_LOG_FLAG",0)
    //ds_handlers["ETO_CONTROL"].HSET("ETO_UPDATE_VALUE",nil)
    return cf.CF_DISABLE   
}
func update_eto_bins( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{
   
   return cf.CF_HALT
}

func make_measuremeent( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{
   
   return cf.CF_DISABLE
}

/*
 def update_eto_bins(self, *parameters):
        #print(int(self.ds_handlers["ETO_CONTROL"].hget("ETO_UPDATE_FLAG")))
        if int(self.ds_handlers["ETO_CONTROL"].hget("ETO_UPDATE_FLAG")) == 1:  ### already done
            return True
            
            
        self.ds_handlers["ETO_CONTROL"].hset("ETO_UPDATE_FLAG",1) # set eto update done 
        
        # find eto with lowest priority
        eto = self.find_eto()
        self.ds_handlers["ETO_CONTROL"].hset("ETO_UPDATE_VALUE",eto)
        if eto ==  None:
           return False
        self.reference_eto = eto
        
        
        
        rain = self.find_rain()
        self.reference_rain = self.find_rain()
        
        for i in self.eto_hash_table.hkeys():  # update eto for irrigation valves
           
           new_value = float(self.eto_hash_table.hget(i)) + float(eto)
           print("eto_update",i,new_value)
           self.eto_hash_table.hset(i,new_value)
           
           
        print("logging sprinkler_data")
        self.log_sprinkler_data()
        return True
*/

/*
    
 def add_eto_chains(eto, cf):

    #
    #
    #  This chain is for diagnostic purposes
    #
    cf.define_chain("test_generator",False)
    cf.insert.enable_chains(["eto_make_measurements"])
    cf.insert.one_step(eto.new_day_rollover)
    cf.insert.enable_chains(["update_eto_bins"])
    cf.insert.terminate()
    

    
    cf.define_chain("enable_measurement", True)
    cf.insert.log("starting enable_measurement")
    
    cf.insert.wait_tod_le(hour=7)  ### Wait till time is less than 8
   
    cf.insert.wait_tod_ge( hour =  7 ) #### Then when time turns 8 act    
    
    cf.insert.one_step(eto.new_day_rollover)
    
    cf.insert.wait_tod_le(hour=8)  ### Wait till time is less than 8
   
    cf.insert.wait_tod_ge( hour =  8 ) #### Then when time turns 8 act
    #
    # We do the day tick here
    #
    #
    
    cf.insert.enable_chains(["eto_make_measurements"])
    cf.insert.log("enabling making_measurement")
    cf.insert.wait_tod_ge(hour=11)
    cf.insert.enable_chains(["update_eto_bins"])
    cf.insert.log("enable_update_eto_bins")
    
    cf.insert.wait_tod_ge( hour =  23 )
    cf.insert.disable_chains(["eto_make_measurements","update_eto_bins"])
    
    cf.insert.wait_event_count( event = "DAY_TICK" )
    cf.insert.reset()

    cf.define_chain("update_eto_bins", False)
    cf.insert.wait_event_count( event = "MINUTE_TICK",count = 8)
    cf.insert.log("updating eto bins")
    cf.insert.wait_function( eto.update_eto_bins )
    cf.insert.disable_chains(["eto_make_measurements"])
    cf.insert.terminate()
 
    
    
    cf.define_chain("eto_make_measurements", False)
    cf.insert.log("starting make measurement")
    
    cf.insert.one_step( eto.make_measurement )

    cf.insert.wait_event_count( event = "MINUTE_TICK",count = 1)
    cf.insert.log("Receiving minute tick")
    
    cf.insert.reset()   
    
    
    
    
}
*/
/*
 * Chain Flow Support functions
 * 
 * 
 */
