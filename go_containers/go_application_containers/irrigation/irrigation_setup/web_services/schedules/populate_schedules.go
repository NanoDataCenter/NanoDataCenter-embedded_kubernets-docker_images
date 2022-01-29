package construct_schedule


import(
    "fmt"
    "encoding/json"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
)






func js_generate_populate_table()string{

  return_value :=  `<script type="text/javascript"> `  
  return_value += web_support.Load_jquery_ajax_components()
  return_value += `
   
 function setup_table(){
    let columns = [  { title:"Select" },{title:"Name"},{title:"Title"} ]

    
    $('#schedule_list').DataTable( {
        pageLength: 50,
        columns: columns
    } );
  }
   function populate_table(){
       let data = {}
       data["master_controller"] = $("#master_server").val()
       data["sub_controller"]    = $("#sub_server").val()
      
       ajax_post_get(ajax_get_schedule , data, ajax_get_function,  "Schedule Data Not Loaded")
       
    }
   function ajax_get_function(data){
      
      set_status_bar("Schedule Data Downloaded")
      let station_data = []
     // let key_entries =Object.keys(data)
     //console.log(key_entries)
     //key_entries.sort()
     //console.log(key_entries)
     //let i = 0
     //for (i = 0;i< key_entries.length;i++){
     //    station_data.push(add_station_entry(eto_resource[key_entries[i]]))
     //}
     //console.log("station data")
     //console.log(station_data)
     let table = $('#schedule_list').DataTable()
     table.clear()
     table.rows.add(station_data)
     table.draw()
      
   }
  
  
  `
  return_value += `</script>`
  return return_value
    
}



func Ajax_add_schedule(input string){  // input master controller, sub_controller, schedule_name , schedule_data
 
    
}
    
func Ajax_delete_schedule(input string){  // input master controller, sub_controller  , schedule_name

}    
    
func Ajax_post_schedules(input map[string]string)string{  // input master controller,sub_controller  output json data
   
   fmt.Println(input) 
    
   schedule_data := make([]map[string]interface{},0)
   bytes,err :=  json.Marshal(schedule_data)
   if err != nil {
       panic(err)
   }
   return string(bytes)
}    
