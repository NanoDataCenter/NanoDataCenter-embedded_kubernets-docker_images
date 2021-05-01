package su // site_utilities


import "lacima.com/site_data"
import "lacima.com/go_setup_containers/site_generation_base/graph_generation/build_configuration"
import "lacima.com/go_setup_containers/site_generation_base/graph_generation/construct_data_structures"
	
var config_file = "/data/redis_server.json"
var site_data_store map[string]interface{}
var Site  string
var Ip    string
var Port  int

var Bc_Rec *bc.Build_Configuration 
var Cd_Rec *cd.Package_Constructor


func Setup_Site_File(){

	site_data_store = get_site_data.Get_site_data(config_file)
    Site = site_data_store["site"].(string)
	Ip   = site_data_store["host"].(string)
	Port = int(site_data_store["port"].(float64))
}

func Setup_graph_generation(){
   bc.Graph_support_init(Ip,Port)
   Bc_Rec = bc.Construct_build_configuration()
   Cd_Rec = cd.Construct_Data_Structures(Bc_Rec)


}

func Done(){

 Bc_Rec.Check_namespace()
 Bc_Rec.Store_keys() 


}