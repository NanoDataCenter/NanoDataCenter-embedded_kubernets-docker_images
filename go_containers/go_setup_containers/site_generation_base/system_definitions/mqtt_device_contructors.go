
package sys_defs
import "lacima.com/go_setup_containers/site_generation_base/site_generation_utilities"



func construct_mqtt_device_defintions() {

  
  su.Bc_Rec.Add_header_node("MQTT_SETUP","mqtt_setup",make(map[string]interface{}))
  su.Cd_Rec.Construct_package("TOPIC_STATUS")
  su.Cd_Rec.Add_hash("TOPIC_STATUS")
  su.Cd_Rec.Add_hash("BAD_TOPIC_STATUS")
  su.Cd_Rec.Add_hash("TIME_STATUS")
  su.Cd_Rec.Create_postgres_stream( "POSTGRES_STREAM","admin","password","admin",30*24*3600)
  su.Cd_Rec.Close_package_construction()
  su.Construct_incident_logging("MQTT_LOG","mqtt_log")
  
  su.Bc_Rec.Add_header_node("MQTT_DEVICES","mqtt_devices",make(map[string]interface{}))
  // generate device device_class
  generate_test_device()
  //
  //
  //
  su.Bc_Rec.End_header_node("MQTT_DEVICES","mqtt_devices")
  su.Bc_Rec.End_header_node("MQTT_SETUP","mqtt_setup")
}

func generate_basic_information(class_name , description string, device_list,topic_list []string , contact_time_seconds int) map[string]interface{}{

  topic_list = append(topic_list,"present")  
    
  test_map := make(map[string]interface{})
  test_map["device_class"]   = class_name
  test_map["description"]    = description
  test_map["device_list"]    = device_list  
  test_map["topic_list"]     = topic_list 
  test_map["contact_time"]   = contact_time_seconds  // for present message
  
  return test_map
    
    
}


func generate_test_device(){
  device_class := "test_mqtt"
  device_list  := []string{"test_device"}
  description  := "this is an example is for testing mqtt interface"
  topic_list   := []string{ "topic_1","topic_2","topic_3"}
  contact_time_seconds := 60 //one minute
  test_map := generate_basic_information(device_class,description,device_list,topic_list,contact_time_seconds)
  su.Bc_Rec.Add_info_node("MQTT_DEVICE","test_device",test_map)
    
    
    
}
