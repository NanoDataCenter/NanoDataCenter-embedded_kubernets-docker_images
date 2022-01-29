package eto_setup



import (
    //"fmt"
    //"strings"
    "encoding/json"
    "net/http"
    "html/template"
    //"github.com/tidwall/gjson"
    //"encoding/json"
    "lacima.com/redis_support/generate_handlers"
	"lacima.com/redis_support/graph_query"
    "lacima.com/redis_support/redis_handlers"
   "lacima.com/go_application_containers/irrigation/irrigation_libraries/irrigation_files_library"  
   
    //"lacima.com/Patterns/web_server_support/jquery_react_support"
    "lacima.com/Patterns/msgpack_2"
    
)

var base_templates        *template.Template
var file_loader           irr_files.Irrigation_File_Manager_Type
var eto_definitions_json  string  
var station_data_json     string
const eto_setup_url_save_link string ="ajax/irrigation/eto/eto_setup_store"

var Eto_Accumulation                  redis_handlers.Redis_Hash_Struct
var Eto_Reserve                       redis_handlers.Redis_Hash_Struct

var Eto_Min_level                     redis_handlers.Redis_Hash_Struct
var Eto_Recharge_Rate                 redis_handlers.Redis_Hash_Struct



func Page_init(input *template.Template){
   
    base_templates = input
    file_loader = irr_files.Initialization()
    setup_eto_handlers()
    resolve_setup_file()
    import_station_valve_data()
    
    
    
}



func Generate_page_setup(w http.ResponseWriter, r *http.Request){
    
    page_template ,_ := base_templates.Clone()
    page_html := generate_eto_setup_html()
    template.Must(page_template.New("application").Parse(page_html))     
    data := make(map[string]interface{})
    data["Title"] = "Setup ETO Resources"
    page_template.ExecuteTemplate(w,"bootstrap", data)
    
}





 


func setup_eto_handlers() {

	search_list                     := []string{"ETO_SETUP:ETO_SETUP","ETO_DATA_STRUCTURES"}
	eto_data_structs                := data_handler.Construct_Data_Structures(&search_list)
	
    Eto_Accumulation                = (*eto_data_structs)["ETO_ACCUMULATION"].(redis_handlers.Redis_Hash_Struct)
	Eto_Reserve                     = (*eto_data_structs)["ETO_RESERVE"].(redis_handlers.Redis_Hash_Struct)
    Eto_Min_level                   = (*eto_data_structs)["ETO_MIN_LEVEL"].(redis_handlers.Redis_Hash_Struct)
    Eto_Recharge_Rate               = (*eto_data_structs)["ETO_RECHARGE_RATE"].(redis_handlers.Redis_Hash_Struct)
    
}



func resolve_setup_file(){
    
   
    eto_definitions_json = validate_eto_file(read_eto_file())
    
    store_eto_file(eto_definitions_json)

}

func read_eto_file()string{
    
  
    
    file_data, _ := file_loader.Read_App_File("eto_site_setup.json")
    return file_data
    
}        
    


func validate_eto_file(input string )string{
    var decode_value map[string]map[string]interface{}
    return_value := input
    if err := json.Unmarshal([]byte(input), &decode_value); err != nil {
        // create an empty element for not valid file data
        marshal_value,_ := json.Marshal(make(map[string]map[string]interface{}))
        return_value = string(marshal_value)
        
    }
        
    
    
    return return_value
    
    
}


func Process_new_eto_setup(raw_input string){
    var decode_value map[string]map[string]interface{}
  
    if err := json.Unmarshal([]byte(raw_input), &decode_value); err != nil {
        panic("bad json")
    }else{
        store_eto_file(raw_input)
        eto_definitions_json = raw_input
    }
    
}

    
func store_eto_file(data string){
    var decode_value map[string]map[string]interface{}
    err := file_loader.Write_App_File("eto_site_setup.json",data)
    if err != true {
        panic("should not happen")
    }
    err1 := json.Unmarshal([]byte(data), &decode_value)
    if err1 != nil {
        panic("should not happen")
    }
    
    trim_hashtable(decode_value,Eto_Accumulation,0.)
    trim_hashtable(decode_value,Eto_Reserve,0.)    
    trim_hashtable(decode_value,Eto_Min_level,.07)
    trim_hashtable(decode_value,Eto_Recharge_Rate,0.)
    update_hashtable(decode_value,Eto_Min_level,"recharge_level")
    update_hashtable(decode_value,Eto_Recharge_Rate,"recharge_rate")
}
    
    
func trim_hashtable( input map[string]map[string]interface{}, table redis_handlers.Redis_Hash_Struct, default_value float64){
    
    pack_value := msg_pack_utils.Pack_float64(default_value)
    // remove keys not present  in input
    for _,key := range table.HKeys() {
       if _,ok := input[key]; ok == false {
           table.HDel(key)
       }
        
    }
    for key,_ := range input {
       if err := table.HExists(key); err == false {
           table.HSet(key,pack_value)
       }
    
    }
}
    
    
    
func  update_hashtable(input map[string]map[string]interface{}, table redis_handlers.Redis_Hash_Struct,field string){
    
    for key,value := range input {
        field_value := value[field].(float64)
        table.HSet(key,msg_pack_utils.Pack_float64(field_value))
    }
}    
    
func import_station_valve_data(){
    station_data := make(map[string][]int64)
    nodes := graph_query.Common_qs_search(&[]string{"IRRIGATION_STATIONS:IRRIGATION_STATIONS","IRRIGATION_STATION"})
   
    for _,node := range nodes {
        //fmt.Println("node",node)
        station_name := graph_query.Convert_json_string(node["name"])
        valve_number   := graph_query.Convert_json_int(node["valve_number"])
        valve_data     := generate_range(valve_number)
        station_data[station_name] = valve_data
    }
    temp,_ := json.Marshal(station_data)
    station_data_json = string(temp)
    //fmt.Println(station_data_json)
   
}


func generate_range( number int64)[]int64{
 
    return_value := make([]int64,number)
    
    for i := int64(0); i< number; i++{
        return_value[i] = i
    }
    return return_value
}
    
    

   
    

    
    
    
