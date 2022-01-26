package irrigation

import (
    "lacima.com/go_setup_containers/site_generation_base/site_generation_utilities"
)

var all_master_actions []string
var all_slave_actions []string




func Add_irrigation_actions(){
    
    
  su.Bc_Rec.Add_header_node("IRRIGATION_ACTIONS","IRRIGATION_ACTIONS",make(map[string]interface{}))
   
   clean_filter_properties := make(map[string]interface{})
   clean_filter_properties["charge_time"]   = 60 //seconds
   clean_filter_properties["back_flush"]    = 30 // seconds
   clean_filter_properties["forward_flush"] = 30 // seconds
   su.Bc_Rec.Add_info_node("IRRIGATION_ACTION","CLEAN_FILTER",clean_filter_properties)

   valve_resistance_properties := make(map[string]interface{})  
   valve_resistance_properties["hold_time"] = 5
   valve_resistance_properties["measure_dt"] = 1
   valve_resistance_properties["meas_number"] = 5
   valve_resistance_properties["next_time"] = 5
   su.Bc_Rec.Add_info_node("IRRIGATION_ACTION","VALVE_RESISTANCE",valve_resistance_properties)
   
   
   valve_leak_properties := make(map[string]interface{})  
   valve_leak_properties["charge_time"] = 300
   valve_leak_properties["wait_time"] = 300

   su.Bc_Rec.Add_info_node("IRRIGATION_ACTION","VALVE_LEAK",valve_leak_properties)
   
   
   
   
   su.Bc_Rec.Add_info_node("IRRIGATION_ACTION","OPEN_MASTER_VALVES",make(map[string]interface{}))
   
   
   
   su.Bc_Rec.Add_info_node("IRRIGATION_ACTION","CLOSE_MASTER_VALVES",make(map[string]interface{}))
   
   
  
   
   su.Bc_Rec.Add_info_node("IRRIGATION_ACTION","IRRIGATION_JOB",make(map[string]interface{}))
   
   
  
   
   su.Bc_Rec.Add_info_node("IRRIGATION_ACTION","NUlL_ACTION_JOB",make(map[string]interface{}))
   
   
  
   
   
   
  su.Bc_Rec.End_header_node("IRRIGATION_ACTIONS","IRRIGATION_ACTIONS")     
  // will need to be revised later  
  all_master_actions = []string{"CLEAN_FILTER","VALVE_RESISTANCE","VALVE_LEAK","OPEN_MASTER_VALVES","CLOSE_MASTER_VALVES","IRRIGATION_JOB","NUlL_ACTION_JOB"}
  all_slave_actions  = []string{"VALVE_RESISTANCE","IRRIGATION_JOB","NUlL_ACTION_JOB"}  
    
    
    
    
    
    
    
    
    
}
