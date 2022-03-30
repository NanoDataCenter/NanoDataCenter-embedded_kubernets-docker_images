package construct_schedule


import(
    //"fmt"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
)



func generate_main_component()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("main_form")

    null_list := make([]string,0)
    return_value.Append_line(web_support.Generate_title("Mange Irrigation Schedules"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_select("Select Master Server","master_server",null_list,null_list))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_select("Select Sub Server","sub_server",null_list,null_list))
   return_value.Append_line(web_support.Generate_space("25"))
    
    values := []string{"null","create","edit","copy","delete"}
    text   := []string{"Null Action","Create Schedule","Edit Schedule","Copy Schedule","Delete Schedule"}
    
    return_value.Append_line(web_support.Generate_select("Select Action","schedule_action",values,text))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_table("List of Schedules","schedule_list"))
    return_value.Append_line("</div>")
    return_value.Append_line(js_generate_top_js())
    
    return return_value

}








func js_generate_top_js()string{

  return_value := 
  ` <script type="text/javascript"> 
  
    
    function main_form_start(){
       hide_all_sections()
       show_section("main_form")
       populate_schedule_list()
    }
  
    function main_form_init(){
      
      master_key = Object.keys(master_sub_server)
      master_key.sort()
      jquery_populate_select('#master_server',master_key,master_key,master_server_change)
      let sub_key  = master_key[0]
      let sub_data = master_sub_server[sub_key]
      sub_data.sort()
      jquery_populate_select("#sub_server",sub_data,sub_data,sub_server_change)
      jquery_initalize_select("#schedule_action",main_menu)
      create_schedule_list_table()
      
    }

    

    function main_menu(event,ui){
       var index
       var choice
       choice = $("#schedule_action").val()
       if( choice == "create"){
   
           start_section("add_schedule")
           
       }
       
       if( choice == "edit"){
          let select_index = find_select_index("Schedule_display_",schedule_data.length)
         if( select_index  == -1){
           alert("no schedule selected")
          }else{
             
              edit_schedule( schedule_data[select_index])
          }
          
           
       }
       if( choice == "copy"){
         let select_index = find_select_index("Schedule_display_",schedule_data.length)
         if( select_index  == -1){
           alert("no schedule selected")
          }else{
             
              copy_schedule_go(select_index)
          }
          
       }
       if( choice == "delete"){
         let select_index = find_select_index("Schedule_display_",schedule_data.length)
         if( select_index  == -1){
           alert("no schedule selected")
          }else{
              let item = schedule_data[select_index]
              let name = item["name"]
              if( confirm("Delete Schedule "+name)== true){
                   let data = {}
                   data["server_key"] =  g_server_key
                   data["name"]     =  name
                  
                   ajax_post_get( ajax_delete_schedule,data, populate_schedule_list, "schedule not deleted")
                  
              }
          }
          
           
       }
       $("#schedule_action")[0].selectedIndex = 0;
              
   }      
   
   function master_server_change(event,ui){
      let sub_key  = $("#master_server").val()
      let sub_data = master_sub_server[sub_key]
      sub_data.sort()
      jquery_populate_select("#sub_server",sub_data,sub_data,null)
      populate_schedule_list()   
   }
    
    
   function sub_server_change(event,ui){
     
   
     populate_schedule_list()
   }
   
   function create_schedule_list_table(){
   
      create_table( "#schedule_list",["Select","Name","Description" ])
   
   
   }
   
  
   
   function populate_schedule_list(){
       let data = {}
       data["server_key"]   = "false~" + $("#master_server").val() +"~"+$("#sub_server").val()
       g_server_key = data["server_key"]
      
       ajax_post_get(ajax_get_schedule , data, ajax_get_function,  "Schedule Data Not Loaded")
       
    }
   function ajax_get_function(data){
      schedule_data  = data
      
      console.log(schedule_data)
      
      schedule_data_map = {}
      set_status_bar("Schedule Data Downloaded")
      let row_data = []
      let i = 0
      for (i = 0;i< schedule_data.length;i++){
         let entry =[]
         let name = schedule_data[i]["name"]
         schedule_data_map[name] = true 
         entry.push(radio_button_element("Schedule_display_"+i))
         
         entry.push(schedule_data[i]["name"])
         entry.push(schedule_data[i]["description"])
         row_data.push(entry)
      }
     console.log("row_data",row_data)
     load_table('#schedule_list', row_data)
      
   }
    
   
   
    </script>`
    
  return return_value
    
    
}
