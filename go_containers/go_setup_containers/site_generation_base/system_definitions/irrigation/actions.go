package irrigation

import (
    "lacima.com/go_setup_containers/site_generation_base/site_generation_utilities"
)



var all_master_actions []string
var all_slave_actions  []string


func Add_irrigation_actions(){
    
    
  su.Bc_Rec.Add_header_node("IRRIGATION_ACTIONS","IRRIGATION_ACTIONS",make(map[string]interface{}))
  
  /* ********************************************************************************************* */
  su.Bc_Rec.Add_header_node("IRRIGATION_ACTIONS","CLEAN_FILTER",make(map[string]interface{}))
  
       clean_filter_properties := make(map[string]interface{})
       clean_filter_properties["charge_time"]   = 60 //seconds
       clean_filter_properties["back_flush"]    = 30 // seconds
       clean_filter_properties["forward_flush"] = 30 // seconds
       su.Bc_Rec.Add_info_node("CLEAN_FILTER","clean_filter_1",clean_filter_properties)
       
  su.Bc_Rec.End_header_node("IRRIGATION_ACTIONS","CLEAN_FILTER")     
   
  
  /* ------------------------------------------------------------------------------------------- */
  
  
  
  
  /* ********************************************************************************************* */
  su.Bc_Rec.Add_header_node("IRRIGATION_ACTIONS","VALVE_RESISTANCE",make(map[string]interface{}))  
   
      valve_resistance_properties := make(map[string]interface{})  
      valve_resistance_properties["hold_time"] = 5
      valve_resistance_properties["measure_dt"] = 1
      valve_resistance_properties["meas_number"] = 5
      valve_resistance_properties["next_time"] = 5
      su.Bc_Rec.Add_info_node("VALVE_RESISTANCE","valve_resistance_1",valve_resistance_properties)
   
  su.Bc_Rec.End_header_node("IRRIGATION_ACTIONS","VALVE_RESISTANCE")
  
  /* ------------------------------------------------------------------------------------------- */
   
   /* ********************************************************************************************* */
  
  su.Bc_Rec.Add_header_node("IRRIGATION_ACTIONS","VALVE_LEAK",make(map[string]interface{}))
  
     valve_leak_properties := make(map[string]interface{})  
     valve_leak_properties["charge_time"] = 300
     valve_leak_properties["wait_time"] = 300
     su.Bc_Rec.Add_info_node("VALVE_LEAK","valve_leak_1",valve_leak_properties)
   
  su.Bc_Rec.End_header_node("IRRIGATION_ACTIONS","VALVE_LEAK")     
   
  /* ------------------------------------------------------------------------------------------- */
  
  
   /* ********************************************************************************************* */
   
  su.Bc_Rec.Add_header_node("IRRIGATION_ACTIONS","OPEN_MASTER_VALVE",make(map[string]interface{}))   
 

     su.Bc_Rec.Add_info_node("OPEN_MASTER_VALVE","open_master_valve_1",make(map[string]interface{}))
 
  su.Bc_Rec.End_header_node("IRRIGATION_ACTIONS","OPEN_MASTER_VALVE")     
   
  /* ------------------------------------------------------------------------------------------- */
  
   /* ********************************************************************************************* */
  su.Bc_Rec.Add_header_node("IRRIGATION_ACTIONS","CLOSE_MASTER_VALVES",make(map[string]interface{}))   
   
    su.Bc_Rec.Add_info_node("CLOSE_MASTER_VALVES","close_master_valve_1",make(map[string]interface{}))

  su.Bc_Rec.End_header_node("IRRIGATION_ACTIONS","CLOSE_MASTER_VALVES")     
  /* ------------------------------------------------------------------------------------------- */ 
   
  
 /* ********************************************************************************************* */  
su.Bc_Rec.Add_header_node("IRRIGATION_ACTIONS","IRRIGATION_SCHEDULE",make(map[string]interface{}))   

   su.Bc_Rec.Add_info_node("IRRIGATION_OPERATIONS","irrigation_operations_1",make(map[string]interface{}))
   
 su.Bc_Rec.End_header_node("IRRIGATION_ACTIONS","IRRIGATION_SCHEDULE")     
  /* ------------------------------------------------------------------------------------------- */
   
 su.Bc_Rec.End_header_node("IRRIGATION_ACTIONS","IRRIGATION_ACTIONS")   
 
 all_master_actions = []string{"CLEAN_FILTER:clean_filter_1","VALVE_LEAK:valve_leak_1","OPEN_MASTER_VALVE:open_master_valve_1","CLOSE_MASTER_VALVES:close_master_valve_1"}
 all_slave_actions  = []string{"IRRIGATION_OPERATIONS:irrigation_operations_1","VALVE_RESISTANCE:valve_resistance_1"}


   
}   
  
   
   
   
 
 
    
  
   
   
   
