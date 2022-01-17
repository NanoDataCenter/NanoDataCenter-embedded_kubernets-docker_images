package eto_adjust

import (
    //"fmt"
    //"strings"
    "net/http"
    "html/template"
    //"encoding/json"
    "lacima.com/redis_support/generate_handlers"
	//"lacima.com/redis_support/graph_query"
    "lacima.com/redis_support/redis_handlers"
  
    "lacima.com/Patterns/web_server_support/jquery_react_support"
    //"lacima.com/Patterns/msgpack_2"
    //"github.com/vmihailenco/msgpack/v5"
)

var base_templates                    *template.Template
var Eto_Accumulation                  redis_handlers.Redis_Hash_Struct
var Eto_Reserve                       redis_handlers.Redis_Hash_Struct

var Eto_Min_level                     redis_handlers.Redis_Hash_Struct
var Eto_Recharge_Rate                 redis_handlers.Redis_Hash_Struct

var eto_adjust_json                   string
var eto_adjust_raw                    map[string]map[string]float64
const eto_setup_url_save_link         string = "ajax/irrigation/eto/eto_adjust_store"


func Page_init(input *template.Template){
    //initialize_eto_data_structures()
    base_templates = input

    setup_eto_handlers()
    
    
}

func setup_eto_handlers() {

	search_list                     := []string{"ETO_SETUP:ETO_SETUP","ETO_DATA_STRUCTURES"}
	eto_data_structs                := data_handler.Construct_Data_Structures(&search_list)
	
    Eto_Accumulation                = (*eto_data_structs)["ETO_ACCUMULATION"].(redis_handlers.Redis_Hash_Struct)
	Eto_Reserve                     = (*eto_data_structs)["ETO_RESERVE"].(redis_handlers.Redis_Hash_Struct)
    Eto_Min_level                   = (*eto_data_structs)["ETO_MIN_LEVEL"].(redis_handlers.Redis_Hash_Struct)
    Eto_Recharge_Rate               = (*eto_data_structs)["ETO_RECHARGE_RATE"].(redis_handlers.Redis_Hash_Struct)
    
}


func Generate_page_adjust(w http.ResponseWriter, r *http.Request){
    
    generate_json_data()
   page_template ,_ := base_templates.Clone()
    page_html := generate_eto_manage()
    template.Must(page_template.New("application").Parse(page_html))        
    data := make(map[string]interface{})
    data["Title"] = "Modify ETO Resources"
    page_template.ExecuteTemplate(w,"bootstrap", data)
    
}

func generate_eto_manage()string{
    return generate_eto_manage_html()+
           generate_eto_manage_js()
}

func generate_eto_manage_html()string{
  html := `
<div class="container">
 <h3 id=edit_panel_header> Make ETO Adjustments </h3>
       <div>
        <input type="button" id = "master_save" value="Save"  data-inline="true"  /> 
        <input type="button" id = "master_reset" value="Reset" data-inline="true"  /> 
       </div>
     
<h3>Mange ETO and ETO Reserve Values</h3>
<h4>Select Select for ETO Settings</h4> 
    <select id="eto_data">
        <option selected value=0>No Action</option>
	        <option value=1>Zero Selected ETO Data</option>
	        <option value=2>Subtract .01 inch from ETO Data</option>
	        <option value=3>Add .01 inch to ETO Data</option>
	        <option value=4>Subtract .05 inch from ETO Data</option>
	        <option value=5>Add .05 inch to ETO Data</option>
	        <option value=6>Select All Elements</option>
	        <option value=7>Unselect All Elememts</option>
	        
        
     </select>
    <h4>Select Manage Reserve Settings</h4> 
    <select id="eto_reserve">
        <option selected value=0>No Action</option>
	        <option value=1>Zero Selected Reserve Data</option>
	        <option value=2>Move .01 inch to ETO Data</option>
	        <option value=3>Move .01 inch to Reserve Data</option>
	        <option value=4>Move .05 inch from ETO Data</option>
	        <option value=5>Move .05 inch to Reserve Data</option>
	        <option value=6>Select All Elements</option>
	        <option value=7>Unselect All Elememts</option>
        
     </select>
</div>
<div style="margin-top:50px"></div>
<div class="container">
     
     <h5>List of ETO Stations</h5>
     <div style="margin-top:10px"></div>
     <table id="resource_list" class="display" width="100%"></table>

     



</div>
</div>
`
return html 
}

func generate_eto_manage_js()string{
   js_data := `
        <script type="text/javascript"> 
        eto_resource_json = '`+eto_adjust_json+`'
        //console.log(eto_resource_json)
        eto_resource      = JSON.parse(eto_resource_json)
        //console.log(eto_resource)
        url_link_save            = '`+eto_setup_url_save_link+`'
        console.log(url_link_save)
        
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
        
         
          
          ajax_post_confirmation(url_link_save, eto_resource,"Do you wish to save eto data?","ETO Data Saved","ETO Data Not Saved") 
        }


        table_columns = [  
             { title: "check" },
             { title: "Key" },
             { title: "ETO inch" },
             { title: "Reserve inch" },
            
          ]   
        
     
        function setup_table(){
  
           $('#resource_list').DataTable( {
                        pageLength: 50,
                        columns: table_columns
                   } );
         }
   
   
function refresh_entries(){
     //console.log(eto_resource)
     let table_data = []
     let key_entries =  Object.keys(eto_resource)
     
     key_entries.sort()
     
     let i = 0
     for (i = 0;i< key_entries.length;i++){
         key = key_entries[i]
         let table_entry = []
         let ref_data    = eto_resource[key]
         let label = " "
         let check_box = '<div class="form-check">\n'+
         '<label class=""btn  btn-toggle" " for="'+key+'">\n'+     	   
         '<input type="checkbox" class="btn  btn-toggle"  id="'+key+'" name="optradio" value='+key+'>'+label+ '</label></div>'
         table_entry.push(check_box)
         table_entry.push(key)
         table_entry.push(ref_data["eto"].toFixed(2))
         table_entry.push(ref_data["reserve"].toFixed(2))
         table_data.push(table_entry)
     }
     
     let table = $('#resource_list').DataTable()
     table.clear()
     table.rows.add(table_data)
     table.draw()
          
  
  }
  function eto_menu(){
      
     

      let keys = Object.keys(eto_resource)
      let choice = $("#eto_data").val()
      
      handle_eto_table_choice(keys,choice)
      
  
    $("#eto_data")[0].selectedIndex = 0
  }
  
  
  function reserve_menu(){
      

      let keys = Object.keys(eto_resource)
      let choice = $("#eto_reserve").val()
      handle_reserve_table_choice(keys,choice)
      
      

  
    $("#eto_reserve")[0].selectedIndex = 0;
  }
  
  
  function handle_eto_table_choice(keys,choice){
    let selected_keys = check_state(keys)
    switch(choice){
        
        case '1':
                 
                 zero_eto_data(keys)
                 refresh_entries()
                 select_values(selected_keys)
                 break
                 
        case '2':
                
                eto_transfer(keys,-.01)
                refresh_entries()
                select_values(selected_keys)
                break
        
        case '3':
                
                eto_transfer(keys,.01)
                refresh_entries()
                select_values(selected_keys)
                break
                
        case '4':
                
                eto_transfer(keys,-.05)
                refresh_entries()
                select_values(selected_keys)
                break
                
        case '5':
                
                eto_transfer(keys,.05)
                refresh_entries()
                select_values(selected_keys)
                break
                
         case '6':
                
                select_values(keys)
                break
            
        case '7':
               unselect_values(keys)
               break
    } 
}   
  
function handle_reserve_table_choice(keys,choice){
    let selected_keys = check_state(keys)
    switch(choice){
        
        
          case '1':
                 zero_reserve_data(keys)
                 refresh_entries()
                 select_values(selected_keys)
                 break
                 
        case '2':
                
                reserve_transfer(keys,-.01)
                refresh_entries()
                select_values(selected_keys)
                break
        
        case '3':
                
                reserve_transfer(keys,.01)
                refresh_entries()
                select_values(selected_keys)
                break
                
        case '4':
                
                reserve_transfer(keys,-.05)
                refresh_entries()
                select_values(selected_keys)
                break
                
        case '5':
                
                reserve_transfer(keys,.05)
                refresh_entries()
                select_values(selected_keys)
                break
        case '6':
                select_values(keys)
                break
            
        case '7':
               unselect_values(keys)
    } 
     
}

function eto_transfer(keys,delta_value ){
    checked_keys = check_state(keys)
    for (i = 0; i<checked_keys.length;i++){
      key = checked_keys[i]
      value = eto_resource[key]["eto"] + delta_value
      if( value < 0.){
          value = 0.
       }
       eto_resource[key]["eto"] = value
        
    }



}
  
function zero_eto_data(keys){
    checked_keys = check_state(keys)
    for (i = 0; i<checked_keys.length;i++){
      key = checked_keys[i]
      
       eto_resource[key]["eto"] = 0
        
    }

}

function reserve_transfer(keys,delta_value){
    checked_keys = check_state(keys)
    for (i = 0; i<checked_keys.length;i++){
      key = checked_keys[i]
      eto     = eto_resource[key]["eto"]
      reserve = eto_resource[key]["reserve"]
      
      if (delta_value < 0) {
          if ( Math.abs(delta_value) > reserve ) {
              delta_value = reserve
          }
          eto = eto + Math.abs(delta_value)
          reserve = reserve - Math.abs(delta_value)
      }else{
        if ( Math.abs(delta_value) > eto ) {
            delta_value = eto
        }
        eto = eto - Math.abs(delta_value)
        reserve = reserve + Math.abs(delta_value)
     }
     eto_resource[key]["eto"] = eto 
     eto_resource[key]["reserve"] = reserve
   }
}
  
function zero_reserve_data(keys){
   checked_keys = check_state(keys)
   for (i = 0; i<checked_keys.length;i++){
       key = checked_keys[i]
       eto_resource[key]["reserve"] = 0.
   }

}
  
  
        
$(document).ready(
 function()
 {  
     
    $("#master_save").bind('click',master_save_control)
   $("#master_reset").bind('click',master_reset_control)
   
   // eto control
   $("#eto_data").bind('change',eto_menu)
   $("#eto_data")[0].selectedIndex = 0
   ;  
   // reseve control
   $("#eto_reserve").bind('change',reserve_menu)
   $("#eto_reserve")[0].selectedIndex = 0;
   setup_table()
   refresh_entries()
 }


)
        
    `    
        
        
    js_data += web_support.Load_jquery_ajax_components()+ web_support.Check_box_state_components()+   `</script>`   
        
    
    return js_data
    
    
}
 
