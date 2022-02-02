package web_support

import (
    //"fmt"
    "strings"
    //"encoding/json"
    
)

type Sub_component_type struct{
  
   Section_name        string
   Html_buffer        []string
   
   
}


type Top_component_type struct{

   global_js          string
   starting_component string
   Section_components map[string]bool
   Html_buffer        []string
   Web_variables      map[string]string
   Ajax_variables     map[string]string
}

func Construct_web_components( top_component, starting_component string,web_variables, ajax_variables map[string]string, global_js string)Top_component_type{
    
     var return_value                 Top_component_type
     return_value.global_js           = global_js
     return_value.starting_component  = starting_component
     return_value.Section_components  = make(map[string]bool)
     
     return_value.Web_variables       = web_variables
     return_value.Ajax_variables      = ajax_variables
     
     
     return_value.Html_buffer         = make([]string,0)
     
     
     return_value.Html_buffer         = append(return_value.Html_buffer,`<div class="container" id="`+top_component+`">`)
     return_value.dump_json_variables()
     return_value.dump_ajax_variables()
     return return_value
}


func (r *Top_component_type)Append_line(input string){
    r.Html_buffer         = append(r.Html_buffer,input+"\n")
}


func (r *Top_component_type)Append_start_script(){
    r.Append_line(`<script type="text/javascript"> ` )
}

func (r *Top_component_type)Append_end_script(){
    r.Append_line(`</script> ` )
}


func (r *Top_component_type)dump_json_variables(){
   
    
    r.Append_start_script()
    for key,value := range r.Web_variables{

       input := key+`_json ='`+value+"'"
       r.Append_line(input)
       input =  key+ "= JSON.parse("+key+"_json)"
       r.Append_line(input) 
    }
    r.Append_end_script()
    
}
func (r *Top_component_type)dump_ajax_variables(){
   
    
    r.Append_start_script()
    for key,value := range r.Ajax_variables{
      
     
       input := key + ` = '`+value+`'`
       r.Append_line(input)
        
    }
    r.Append_end_script()
    
}


func (r *Top_component_type)add_section_control(){
       
       r.Append_start_script()
       r.Append_line("function hide_all_sections(){")
       
       for section,_ := range r.Section_components{
            
            r.Append_line(`$("#`+section+`").hide()`)
       }
       r.Append_line("}")
       r.Append_line("function show_section(section){")

       for section,_ := range r.Section_components{
            r.Append_line(`if ( section =="`+section+`"){`)
        
            r.Append_line(`$("#`+section+`").show()`)
            r.Append_line(`}`)
       }
       r.Append_line("}")
       
       r.Append_line("function hide_section(section){")

       for section,_ := range r.Section_components{
            r.Append_line(`if ( section =="`+section+`"){`)
        
            r.Append_line(`$("#`+section+`").hide()`)
            r.Append_line(`}`)
       }
       r.Append_line("}")
       r.Append_line("function init_sections(){")
       for section,_ := range r.Section_components{
            r.Append_line(section+`_init()`)
       }
       
       r.Append_line("}")       
       
       r.Append_line("function start_section(section){")
      
       for section,_ := range r.Section_components{
            r.Append_line(`if ( section =="`+section+`"){`)
           
            r.Append_line(`  `+section+`_start()`)
            r.Append_line(`}`)
       }
     
       r.Append_line("}")   
      
      r.Append_end_script()
     
    
}


func (r *Top_component_type )Generate_ending()string{
   r.add_section_control()
   r.Append_start_script()
   r.Append_line(r.global_js)
   r.Append_end_script()
   r.Html_buffer         = append(r.Html_buffer,`</div>`) 
   
   return strings.Join(r.Html_buffer,"\n")
    
    
}

func (r *Top_component_type)Transfer_buffer( input []string){
    for _,value := range input {
      r.Append_line(value)      
    }
}

func (r *Top_component_type)Add_section(input Sub_component_type){
   
   if  _,ok := r.Section_components[input.Section_name] ; ok == true{
       panic("duplicate section")
   }
   r.Section_components[input.Section_name] = true
   r.Transfer_buffer(input.Html_buffer)
}



func Construct_subsystem(section_name string)Sub_component_type{
    
   var return_value Sub_component_type
   return_value.Section_name = section_name
   return_value.Html_buffer  = make([]string,0)
   return_value.Html_buffer         = append(return_value.Html_buffer,`<div class="container" id="`+section_name+`">`) 
   return return_value 
}






func  (r *Sub_component_type)Dump_buffer()string{
    
    return strings.Join(r.Html_buffer,"\n")
}

func (r *Sub_component_type)Append_line(input string){
    r.Html_buffer         = append(r.Html_buffer,input+"\n")
}


func (r *Sub_component_type)Append_start_script(){
    r.Append_line(`<script type="text/javascript"> ` )
}

func (r *Sub_component_type)Append_end_script(){
    r.Append_line(`</script>` )
}
