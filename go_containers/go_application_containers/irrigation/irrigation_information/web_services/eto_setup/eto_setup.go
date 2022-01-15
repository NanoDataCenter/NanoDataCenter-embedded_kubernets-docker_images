package eto_setup

import (
   
	"strings"
    
	//"lacima.com/cf_control"
	//"lacima.com/go_application_containers/irrigation/eto/eto_calc/eto_support"
	//"lacima.com/redis_support/generate_handlers"
	//"lacima.com/redis_support/graph_query"
	//"lacima.com/redis_support/redis_handlers"
	//"time"
    //"lacima.com/Patterns/msgpack_2"
    //"lacima.com/server_libraries/postgres"
    //"lacima.com/go_application_containers/irrigation/irrigation_information/web_services/eto_setup/eto_setup"
)





func generate_eto_setup_html()string{
    html_array := make([]string,13)
    html_array[0] = include_style()
    html_array[1] = include_top_js()
    html_array[2] = include_station_javascript()
    html_array[3] = `<div class="container"></div>`
    html_array[4] = generate_main_html("main_div")

    html_array[5] = create_change_recharge_level("recharge_div")
    html_array[6] = create_crop_utilization("crop_utilization_div")
    html_array[7] = create_change_salt_flush("salt_flush_div")
    html_array[8] = create_change_sprayer_efficiency("sprayer_efficiency_div")
    html_array[9] = create_change_sprayer_rate("sprayer_rate_div")
    html_array[10] = create_change_tree_radius("tree_radius_div")
    html_array[11] = create_new_station_html("new_station_div")
    html_array[12] = `</div>`
    return strings.Join(html_array,"\n")
}

func include_style()string{
    
return_value := `
 <style>


.auto {
  -ms-touch-action: auto;
  touch-action: auto;
} 
</style>
`
return return_value
 
}



func include_top_js()string{
    js_data := `
        <script type="text/javascript"> 
        eto_resource_json = '`+eto_definitions_json+`'
        eto_resource      = new Map(Object.entries(JSON.parse(eto_resource_json)))
        station_config_data_json = '`+station_data_json+`'
        station_config_data = JSON.parse(station_config_data_json)
        
        
       function zeroPad(num, places) {
           var zero = places - num.toString().length + 1;
           return Array(+(zero > 0 && zero)).join("0") + num;
        }

        
        function add_new_station( new_station ){
           let key = new_station["station"]+"__"+zeroPad(new_station["valve"],2)
          
           
           if (eto_resource.has(key) != true){
 
               new_station["key"] = key 
               new_station = calculate_eto_recharge_rate(new_station)
               //console.log("new_station")
               //console.log(new_station)
               eto_resource.set(key,new_station)
               //console.log(eto_resource)
               refresh_entries(eto_resource)
               return true
           }
           alert("duplicate_key")
           return false
           
         }
        
        
        
        
        
        function setup_main_screen(){
            $("#main_div").show()
            $("#new_station_div").hide()
            $("#recharge_div").hide()
            $("#crop_utilization_div").hide()
            $("#salt_flush_div").hide()
            $("#sprayer_efficiency_div").hide()
            $("#sprayer_rate_div").hide()
            $("#tree_radius_div").hide()
         }
         
         function setup_aux_screen( div ){
            $("#main_div").hide()
            div.show()
         }
         
         var eto_ref_map     = new Map();
         
         eto_ref_map["recharge_level"]       =  .07
         eto_ref_map["crop_utilization"]     =  .8
         eto_ref_map["sprayer_effiency"]     =  .8
         eto_ref_map["salt_flush"]           =  .1,
         eto_ref_map["sprayer_rate"]         =  14.5
         eto_ref_map["tree_radius"]          =  6.0
         
         
         var eto_holding_map = new Map();
         
         function update_values(field,value){
            //console.log(field,value)
            let update_keys = check_state()
            let i=0
            for (i=0;i<update_keys.length;i++){
                key = update_keys[i]
                data = eto_resource.get(key)
                data[field] = value
                data = calculate_eto_recharge_rate(data)
                eto_resource.set(key,data)
               
            }
             refresh_entries(eto_resource)
             select_values(update_keys)
          }
         
         
        </script>`
        
    return js_data
}     



func generate_main_html(div_name string)string{
    html := `
     <div id="`+div_name+`"  class="container">
     
      <div id="master_save_div">
      <h3 id=edit_panel_header> Make ETO modifications </h3>
       <div>
        <input type="button" id = "master_save" value="Save"  data-inline="true"  /> 
        <input type="button" id = "master_reset" value="Reset" data-inline="true"  /> 
       </div>
     </div>
    <h5>Select Action</h5>
    <select name="#action-choice" id="action-choice">
         <option   value="nop">No Operation</option>
         <option   value="recharge">Recharge Level</option>
         <option   value="crop_util">Crop Utilization</option>
         <option   value="salt">Salt Flush</option>
         <option   value="sprayer_efficiency">Sprayer Efficiency</option>
         <option   value="sprayer_rate">Sprayer Rate</option>
         <option   value="tree_radius">Tree Radius</option>
         <option   value="add">Add Eto Resource</option>
         <option   value="delete">Delete Eto Resource</option>
         <option   value="select_all">Select All Valves</option>
         <option   value="unselect_all">Unselect All Valves</option>
     </select>
     <div style="margin-top:25px"></div>
     <h5>List of ETO Stations</h5>
     <div style="margin-top:10px"></div>
     <table id="resource_list" class="display" width="100%"></table>

     


   </div>
 `
   return  html+"\n"+generate_main_js()
    
    
}

func generate_main_js()string{

    js := `
    
<script type="text/javascript">    

function master_reset_control()
{
       var result = confirm("Do you wish to reset data?");  
      
       if( result == true )
       {
          window.location = window.location

       }
}

function master_save_control()
{
       var result = confirm("Do you wish to save data?");  
      
       if( result == true )
       {
          setup_main_screen()

       }
}


function main_menu(event,ui){
   var index
   var choice

   choice = $("#action-choice").val()

   if( choice == "recharge")
   {
    
       let update_keys = check_state()
       if (update_keys.length == 0){
         alert("no valve selected")
         return
       }else{
       let key = update_keys[0]
       let data = eto_resource.get(key)
       setup_aux_screen( $("#recharge_div"))
       recharge_div_open(data["recharge_level"])
       }
   }
   if( choice == "crop_util")
   {
       let update_keys = check_state()
       if (update_keys.length == 0){
         alert("no valve selected")
         
       }else{
       let key = update_keys[0]
        let data = eto_resource.get(key)
       setup_aux_screen( $("#crop_utilization_div"))
       crop_utilization_div_open(data["crop_utilization"])
       }
   }
   
   if( choice == "salt")
   {
       
       let update_keys = check_state()
       if (update_keys.length == 0){
         alert("no valve selected")
         return
       }else{
       let key = update_keys[0]
        let data = eto_resource.get(key)
       setup_aux_screen( $("#salt_flush_div"))
       salt_flush_div_open(data["salt_flush"])
      }
   }

   
   if( choice == "sprayer_efficiency")
   {
       let update_keys = check_state()
       if (update_keys.length == 0){
         alert("no valve selected")
         
       }else{
       let key = update_keys[0]
        let data = eto_resource.get(key)
       setup_aux_screen( $("#sprayer_efficiency_div"))
       sprayer_efficiency_div_open(data["sprayer_effiency"])
       }
   }

   if( choice == "sprayer_rate")
   {
       let update_keys = check_state()
       if (update_keys.length == 0){
         alert("no valve selected")
        
       }else{
       let key = update_keys[0]
        let data = eto_resource.get(key)
       setup_aux_screen( $("#sprayer_rate_div"))
       sprayer_rate_div_open(data["sprayer_rate"])
       }
   }
   if( choice == "tree_radius")
   {
       let update_keys = check_state()
       if (update_keys.length == 0){
         alert("no valve selected")
    
       }else{
       let key = update_keys[0]
        let data = eto_resource.get(key)
       setup_aux_screen( $("#tree_radius_div"))
       tree_radius_div_open(data["tree_radius"])
       }
   }
   
   
 
   if( choice == "add")
   {
       setup_aux_screen( $("#new_station_div"))
       recharge_div_add_open(eto_ref_map["recharge_level"])
       salt_flush_div_add_open(eto_ref_map["salt_flush"])
       crop_utilization_div_add_open(eto_ref_map["crop_utilization"])
       sprayer_efficiency_div_add_open(eto_ref_map["sprayer_effiency"])
       sprayer_rate_div_add_open(eto_ref_map["sprayer_rate"])
       tree_radius_div_add_open(eto_ref_map["tree_radius"])
       
       
      
   }
 
   if( choice == "delete"){
   
       let update_keys = check_state()
       if (update_keys.length == 0){
         alert("no valve selected")
         
       }else{
       
      let result = confirm("Do you wish to delete selected valves?");  
      
       if( result == true )
       {
          for (i=0;i<update_keys.length;i++){
              
              eto_resource.delete(update_keys[i])
          }
          refresh_entries(eto_resource)
          

       }// confirm
       unselect_all()
     }//else delete
       
   }// if delete
   
   if( choice == "select_all"){
       
       select_all()
   }
   if( choice == "unselect_all"){
       
       unselect_all()
   }
    
  
   $("#action-choice")[0].selectedIndex = 0;
   
}

$(document).ready(
function()
{
   setup_main_screen()
   setup_table()
   recharge_div_load_controls()
   crop_utilization_div_load_controls()
   salt_flush_div_load_controls()
   sprayer_efficiency_div_load_controls()
   sprayer_rate_div_load_controls()
   tree_radius_div_load_controls()
   setup_station_controls()
   
   
   $("#master_save").bind('click',master_save_control)
   $("#master_reset").bind('click',master_reset_control)
   
   
   $("#action-choice").bind('change',main_menu)
   $("#action-choice")[0].selectedIndex = 0;
})
    </script>
`
  return js

}




func include_station_javascript()string{

  js:= `
<script type="text/javascript">
  var station_keys = [ "key","station","valve","recharge_rate","recharge_level" ,"salt_flush", "crop_utilization","sprayer_effiency","sprayer_rate","tree_radius"]
  var station_data = []
  
  table_columns = [  { title: "Select"},
             { title: "Key" },
             { title: "Station" },
             { title: "Valve" },
             { title: "Recharge Rate" },
             { title: "Recharge Level" },
             { title: "Salt Flush"},
             { title: "Crop Utilization" },
             { title: "Sprayer Efficency" },
             { title: "Srayer Rate" },
             { title: "Tree Radius" },
          ]   
             
             
             
  function setup_table(){
  
    $('#resource_list').DataTable( {
        pageLength: 50,
        station_data: station_data ,
        columns: table_columns
    } );
  }
  
  function refresh_entries(eto_resource){
     //console.log(eto_resource)
     station_data = []
     let key_entries =Array.from(  eto_resource.keys())
     //console.log(key_entries)
     key_entries.sort()
     //console.log(key_entries)
     let i = 0
     for (i = 0;i< key_entries.length;i++){
         station_data.push(add_station_entry(eto_resource.get(key_entries[i])))
     }
     //console.log("station data")
     //console.log(station_data)
     let table = $('#resource_list').DataTable()
     table.clear()
     table.rows.add(station_data)
     table.draw()
          
  
  }
  
  
  function check_state(){
   keys = Array.from( eto_resource.keys())
   let check_status = [];
   //console.log(keys)
   for (i= 0;i<keys.length;i++){
      let key = keys[i]
      //console.log("loop")
      //console.log(i)
      //console.log(key)
      if( $("#"+key).is(":checked") == true )
      {       
	         check_status.push(key);
	         
       }
        
    }
   return check_status
     
  }
 
  function add_station_entry(data){
    //console.log(data)
    //let keys = Array.from( data.keys())
    
    let key = data["key"]
    return_value =  []
    label = " "
    check_box = '<div class="form-check">\n'+
    '<label class=""btn  btn-toggle" " for="'+key+'">\n'+     	   
    '<input type="checkbox" class="btn  btn-toggle"  id="'+key+'" name="optradio" value='+key+'>'+label+ '</label></div>'
    
    
    return_value.push(check_box)        

    return_value.push(data["key"])
    return_value.push(data["station"])
    return_value.push(data["valve"])
    return_value.push(data["recharge_rate"].toFixed(2))
    return_value.push(data["recharge_level"].toFixed(2))
    return_value.push(data["salt_flush"].toFixed(2))
    return_value.push(data["crop_utilization"].toFixed(2))
    return_value.push(data["sprayer_effiency"].toFixed(2))
    return_value.push(data["sprayer_rate"].toFixed(2))
    return_value.push(data["tree_radius"].toFixed(2))
    //console.log("return_value")
    //console.log(return_value)
    return return_value
}
 
 
 
 function calculate_eto_recharge_rate( new_entry ){    
   
   let sprayer_rate          = new_entry["sprayer_rate"]
   let tree_radius           = new_entry["tree_radius"]
   let sprinkler_efficiency  = new_entry["sprayer_effiency"]
   let salt_flush            = new_entry["salt_flush"]
   let crop_utilization      = new_entry["crop_utilization"]
   // effective volume is ft3/hr
   // sprayer rate is in Gallons/hr
   // 1 gallon is 0.133681 ft3
   //console.log(sprayer_rate)
   //console.log(tree_radius)
   //console.log(sprinkler_efficiency)
   //console.log(salt_flush)
   //console.log(crop_utilization)
        
   effective_rate = sprayer_rate*sprinkler_efficiency/crop_utilization/(1+salt_flush)
   effective_volume = 0.133681 * effective_rate
   effective_area =  tree_radius*tree_radius*3.14159
   recharge_rate = (effective_volume/effective_area)*12; // converting the eto to inches
   new_entry["recharge_rate"] = recharge_rate
   return new_entry
   
   
   
} 
 
function select_all(){
 let key_entries =Array.from(  eto_resource.keys())
 select_values(key_entries)
}
 
 
function unselect_all(){
 let key_entries =Array.from(  eto_resource.keys())
 unselect_values(key_entries)
}
 
function select_values(keys){

  let i = 0
  for( i= 0;i<keys.length;i++){
      let key = keys[i]
      $("#"+key).prop('checked', true)
   }

}

function unselect_values(keys){
  let i = 0
  for( i= 0;i<keys.length;i++){
      let key = keys[i]
      $("#"+key).prop('checked', false)
   }


}
 
 </script>
 `
 return js
}


 

 
 
 
 
 



    
    
    
