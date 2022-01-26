package construct_schedule


import(
    //"fmt"
    "encoding/json"
    "strings"
    "net/http"
    "html/template"
    "lacima.com/redis_support/graph_query"
)

var base_templates                    *template.Template

var master_table_list                 map[string][]string
var master_table_list_json            string

var valve_list                        map[string]map[string]interface{}
var valve_list_json                   string

var schedule_data                    map[string]map[string]map[string]interface{}
var schedule_json                    string


func Page_init(input *template.Template){
    
    
    construct_master_server_list()
    get_schedule_data()
    base_templates = input

    
    
    
}      


func get_schedule_data(){
    schedule_data = make(map[string]map[string]map[string]interface{})
    temp1,_ := json.Marshal(schedule_data)
    schedule_json = string(temp1)
}
    
func construct_master_server_list(){

    master_table_list = make(map[string][]string)
    valve_list        = make(map[string]map[string]interface{})
    nodes := graph_query.Common_qs_search(&[]string{"IRRIGATION_SERVERS:IRRIGATION_SERVERS","IRRIGATION_SERVER"})
    
    for _,node := range nodes {
       name     := graph_query.Convert_json_string(node["name"])
       master_table_list[name],valve_list[name] = find_subnodes( name )
        
        
    }
    temp,_ := json.Marshal(master_table_list)
    master_table_list_json = string(temp)
    temp1,_ := json.Marshal(valve_list)
    valve_list_json = string(temp1)
    
   
    
}
    
func find_subnodes( master_node string )([]string,map[string]interface{}){
    return_value2 := make(map[string]interface{})
    return_value1 := make([]string,0)
    sub_nodes := graph_query.Common_qs_search(&[]string{"IRRIGATION_SERVERS:IRRIGATION_SERVERS","IRRIGATION_SERVER:"+master_node,"IRRIGATION_SUBSERVER"})
    if len(sub_nodes) == 0{
        return return_value1,return_value2
    }
    for _,sub_node := range(sub_nodes){
        name     := graph_query.Convert_json_string(sub_node["name"])
        
        byte_array := []byte(sub_node["supported_stations"])
        var data map[string][]int
        if err := json.Unmarshal(byte_array, &data); err != nil {
           panic(err)
        }
        
        return_value2[name]=data
        return_value1 = append(return_value1,name)
    }
   
    return return_value1,return_value2


}    



func Generate_page(w http.ResponseWriter, r *http.Request){
    
     
    page_template ,_ := base_templates.Clone()
    page_html := generate_html_js()
    template.Must(page_template.New("application").Parse(page_html))        
    data := make(map[string]interface{})
    data["Title"] = "Modify ETO Resources"
    page_template.ExecuteTemplate(w,"bootstrap", data)
    
}


func generate_html_js()string{
    
    
    return generate_html()+generate_js()
}


func generate_html()string{
   html_array := make([]string,4)
   html_array[0] = generate_top_html()
   html_array[1] = generate_construction_table_html()
   html_array[2] = generate_valve_table_html()
   html_array[3] = generate_table_entry_html()
   return strings.Join(html_array,"\n")
    
    
}

func generate_js()string{
   js_array := make([]string,5)
   js_array[0] = js_generate_global_js()
   js_array[1] = js_generate_top_js()
   js_array[2] = js_generate_construction_table_js()
   js_array[3] = js_generate_valve_table_js()
   js_array[4] = js_generate_table_entry_js()
   return strings.Join(js_array,"\n")    
}
   
func js_generate_global_js()string{
  return_value := 
    ` <script type="text/javascript"> 
    master_sub_server_json ='`+master_table_list_json+`'
    master_sub_server = JSON.parse(master_sub_server_json)

    valve_list_json ='`+valve_list_json+`'
    valve_list = JSON.parse(valve_list_json)
    
    schedule_json ='`+schedule_json+`'
    schedule = JSON.parse(schedule_json)
    
    
    function hide_all_sections(){
    
        $("#main_section").hide()
        $("#table_construction").hide()
        $("#valve_section").hide()
        $("#table_name_section").hide()
    
    }
    
    $(document).ready(
    function()
    {  
       hide_all_sections()
       initialize_main_panel()
       initialize_schedule_construction_panel()
       initialize_valve_construction_panel()
       initialize_table_name_panel()
       start_main_panel()
    
    })
    
    
    </script>`
    
  return return_value
    
    
}
  
    
func generate_top_html()string{
    
  return_value :=
  `
  <div class="container" id="main_section">
 
     
    <h3>Mange Irrigation Schedules</h3>
    <h4>Select Master Server</h4> 
    <select id="master_server">
     </select>
    <h4>Select Sub Server</h4> 
    <select id="sub_server">
    </select>
    <div style="margin-top:20px"></div>
    <h4>Select Select Action</h4> 
    <select id="schedule_action">
    <option value="null">Null Action</option>
    <option value="create">Create Schedule</option>
    <option value="edit">Edit Schedule</option>
    <option value="copy">Copy Schedule</option>
    <option value="delete">Delete Schedule</option>
    </select>
    <div style="margin-top:20px"></div>
   
     <div style="margin-top:20px"></div>
     <h4>List of Schedules</h4>
     <div style="margin-top:20px"></div>
    
     <table id="schedule_list" class="display" width="100%"></table>
    
    
    </div>
    
    
`
 return return_value
}


func js_generate_top_js()string{

  return_value := 
    ` <script type="text/javascript"> 
    function initialize_main_panel(){

      populate_master_select()
      
      attach_action_handler()
      populate_table()

    }
       
    function start_main_panel(){
       hide_all_sections()
       $("#main_section").show()
    }
    
    // supporting function
    
    function populate_master_select(){
      
      master_key = Object.keys(master_sub_server)
      master_key.sort()
      
      for(let i=0; i<master_key.length; i++){
        $('#master_server').append($('<option>').val(master_key[i]).text(master_key[i]));
      }
      let sub_key  = master_key[0]
      let sub_data = master_sub_server[sub_key]
      populate_sub_server_select(sub_data)
      $("#master_server").bind('change',master_change)
     
    }
    
    
    function populate_sub_server_select(sub_select_list){
    
        
        for(let i=0; i<sub_select_list.length; i++){
           $('#sub_server').append($('<option>').val(sub_select_list[i]).text(sub_select_list[i]));
        }
        $("#sub_server")[0].selectedIndex = 0;
    }
    function attach_action_handler(){
    
      $("#schedule_action").bind('change',main_menu)
      $("#schedule_action")[0].selectedIndex = 0;
    }
    
    function populate_table(){
       
       console.log("populate table")
    }
    
    function main_menu(event,ui){
       var index
       var choice
       choice = $("#schedule_action").val()
       if( choice == "create"){
   
           start_table_entry_panel()
           
       }
       
       if( choice == "edit"){
           
           alert("edit")
           
       }
       if( choice == "copy"){
           
           alert("copy")
           
       }
       if( choice == "delete"){
           
           alert("delete")
           
       }
       $("#schedule_action")[0].selectedIndex = 0;
              
   }      
   
   function master_change(event,ui){
      let sub_key  = $("#master_server").val()
      let sub_data = master_sub_server[sub_key]
      populate_sub_server_select(sub_data)
      
   
   }
    

    </script>`
    
  return return_value
    
    
}

func generate_construction_table_html()string{
    
  return_value :=
  `
  <div class="container" id="table_construction">
 
   <h3>Mange Schedules</h3>
    <h4>Master</h4>
    <h4>Slave</h4>
    <h4>Schedule</h4>
    
    
    <h3 > Make Schedule Adjustments </h3>
       <div>
        <input type="button" id = "schedule_save" value="Save"  data-inline="true"  /> 
        <input type="button" id = "schedule_reset" value="Reset" data-inline="true"  /> 
       </div>
    <select id="schedule_action">
    
    </select>
    <div class="container">
     
     <h5>List of Steps</h5>
     <div style="margin-top:10px"></div>
     <table id="step_list" class="display" width="100%"></table>
    </div>
    </div>
    </div>
    
`
 return return_value
}

func js_generate_construction_table_js()string{

  return_value := 
    ` <script type="text/javascript"> 
     
       function initialize_schedule_construction_panel(){
       
       }
       
       function start_construct_table(){
       
       
       }
    </script>`
    
  return return_value
    
    
}   

    

func generate_valve_table_html()string{
    
  return_value :=
  `
  <div class="container" id="valve_section">
 
     
    <h3>Mange Valves</h3>
    <h4>Master</h4>
    <h4>Slave</h4>
    <h4>Schedule</h4>
    <h4>Step</h4>
    
    <h3 > Make Valve Adjustments </h3>
       <div>
        <input type="button" id = "valve_save" value="Save"  data-inline="true"  /> 
        <input type="button" id = "valve_reset" value="Reset" data-inline="true"  /> 
       </div>
     
    <div class="container">
     
     <h5>List of Valves</h5>
     <div style="margin-top:10px"></div>
     <table id="valve_list" class="display" width="100%"></table>
    </div>
    
    </div>
    
    </div>
    
`
 return return_value
}

func js_generate_valve_table_js()string{

  return_value := 
    ` <script type="text/javascript"> 
      
       function initialize_valve_construction_panel(){
       
       
       }
       
       function start_valve_contruction_panel(){
       
       
       }
       
    </script>`
    
  return return_value
    
    
}
   
func generate_table_entry_html()string{
    
  return_value :=
  `
  <div class="container" id="table_name_section">
 
     
    <h3>Enter New Table Name</h3>
    
    
       <div>
        <input type="button" id = "table_name_save" value="Save"  data-inline="true"  /> 
        <input type="button" id = "table_name_abort" value="Abort" data-inline="true"  /> 
       </div>
     <input type="text" id="new_schedule_input">
    
    
    </div>
    
`
 return return_value
}

func js_generate_table_entry_js()string{
  return_value := 
    ` <script type="text/javascript"> 
      
       function initialize_table_name_panel(){
       
       
       }
       function start_table_entry_panel(){
          hide_all_sections()
          $("#table_name_section").show()
       }
    </script>`
    
  return return_value
 
    
}
/*
 var str = $("#myInput").val();
        alert(str);
*/
