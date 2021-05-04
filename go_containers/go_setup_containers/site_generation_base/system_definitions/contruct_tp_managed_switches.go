package sys_defs

import "lacima.com/go_setup_containers/site_generation_base/site_generation_utilities"

func Construct_monitored_switches(){ 
   
    properties := make(map[string]interface{})
    properties["ip"] = "192.168.1.45"
 
    su.Bc_Rec.Add_header_node("TP_SWITCH","switch_office",properties)
	su.Construct_incident_logging("switch_office")
    su.Bc_Rec.End_header_node("TP_SWITCH","switch_office")

    properties = make(map[string]interface{})
    properties["ip"] = "192.168.1.56"
    su.Bc_Rec.Add_header_node("TP_SWITCH","switch_garage",properties)
    su.Construct_incident_logging("switch_garage")
    su.Bc_Rec.End_header_node("TP_SWITCH","switch_garage")
    
}    