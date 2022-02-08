package web_support

import (
    
   "strings"   
    
)


func Generate_select(title,id_tag string,option_values,option_texts []string)string{
    
    
    if len(option_values) != len(option_texts){
        panic("Length not equal")
    }
    last_index := len(option_values) +3-1
    return_value := make([]string,len(option_values)+3)


    return_value[0]  = `<h4>`+title+`</h4>`
    return_value[1] = `<select id="`+id_tag+ `">`
    for index,value := range option_values{
        text := option_texts[index]
        return_value[index+2] =   `<option value="`+value+ `">`+text+`</option>` 
        
    }
    return_value[last_index] = `</select>`
    
    return strings.Join(return_value,"")

}


func Generate_table(title,id_tag string)string{
    
    return_value := make([]string,2)
    return_value[0] =`<h4>`+title+`</h4>`
    return_value[1] = `<table id="`+id_tag+`" class="display" width="100%"></table>`
    return strings.Join(return_value,"")
}


func Generate_title(label string)string{
 
    return "<h3>"+label+"</h3>"
    
}

func Generate_sub_title(id string, label string)string{
 
    return `<h4 id ="`+id+`">`+label+"</h4>"
    
}


func Generate_button(label,id  string)string{
    
     return ` <input type="button" id="`+id+`" value="`+label+`"  data-inline="true"  />`
}

func Generate_input(label,id string)string{
     return_value := make([]string,2)
     return_value[0] = `<h4>`+label+`</h4>`
     return_value[1] = `<input type="text" id="`+id+`">`
     return strings.Join(return_value,"\n")
}


func Generate_space(spacing string)string{
   return `<div style="margin-top:`+spacing+`px"></div>`
}

func Jquery_components_js()string{
    js := `
    function jquery_populate_select(id_tag,value_array,text_array,change_function){
    
        $(id_tag).empty()
        for(let i=0; i<value_array.length; i++){
           $(id_tag).append($('<option>').val(value_array[i]).text(text_array[i]));
        }
        jquery_initalize_select(id_tag,change_function)
    }
    
    function jquery_initalize_select(id_tag,change_function){
        
        $(id_tag)[0].selectedIndex = 0;
        if(change_function != null ){
            $(id_tag).bind('change', change_function)
        }

    }
    
    
    
    function attach_button_handler(id,handler){
      if( handler == null ){
          return
      }
      $(id).bind("click",handler)    
    }
    
    
    function create_table( tag,columns){
      let table_columns = []
      for( let i= 0;i<columns.length;i++){
        table_columns.push( { title:columns[i] } )
      }  
      
      
      dataSet = [] 
      $(tag).DataTable( {
        data: dataSet,
        pageLength: 50,
        columns: table_columns
    } );
  }
  function load_table(tag, data){
  
     let table = $(tag).DataTable()
     table.clear()
     table.rows.add(data)
     table.draw()
  
  
 }
    
    
  function check_box_element(key){
  
        let label = " "
        check_box = '<div class="form-check">\n'+
       '<label class=""btn  btn-toggle" " for="'+key+'">\n'+     	   
       '<input type="checkbox" class="btn  btn-toggle" id="'+key+'" name="optradio" value='+key+'>'+label+ '</label></div>'
        return check_box
  }
  
  function radio_button_element(key){
  
       let label = " "
       radio_button = '<div class="form-check">\n'+
       '<label class=""btn  btn-toggle" " for="'+key+'">\n'+     	   

       '<input type="radio" class="btn  btn-toggle"  id="'+key+'" name="optradio" value='+key+'>'+label+ '</label></div>'
       return radio_button
   }
  
   function find_select_index(base_id,length){
       var i;
       for( i=0; i< length; i++)
       {
         if( $("#"+base_id+i).is(':checked') == true )
         {
           return i;
         }
      }
      return  -1;  // no item selected
     }
     
     function find_check_box_elements(base_id,length){
       var return_value = []
       for( i=0; i< length; i++)
       {
         if( $("#"+base_id+i).is(':checked') == true )
         {
           return_value.push(i)
         }
      }
      return return_value
     
     
     }
    
    function deepcopy(input){
       return JSON.parse(JSON.stringify(input))
   }
    
    function deepslice( input, index,number){
         return_value = []
         if( input.length < index+number){
             alert("bad slice values")
             return return_value
          }
          let i = 0
          for( i = 0; i<input.length;i++){
              if( (i < index)||(i>index+number-1)){
                  return_value.push(deepcopy(input[i]))
               }
          }
          return return_value
    }
    
    function Keys(input){
      return Object.keys(input)
    }
    
    function calculate_move(length,move_position,move_array){
      let return_value = []
      let ref_array    = []
     
      if(move_position+ move_array.length > length){
          move_position = length-move_array.length
       }
    
       for(let i=0;i<length;i++){
         return_value.push(-1)
         ref_array.push(i)
       }
       
       for(let i=0;i<move_array.length;i++){
          
          ref_array[move_array[i]] = -1
          return_value[move_position] = move_array[i]
          move_position += 1
        }
        
        for(let i= 0;i<ref_array.length;i++){
            let temp = ref_array[i]
            
            if( temp == -1){
                continue
             }
             return_valve = inner_move(temp,return_value)
             
         }
        
        return return_value
     }
     
     function inner_move(value,return_value){
        for(let j=0;j<return_value.length;j++){
              if( return_value[j] == -1 ){
                return_value[j] =value
                return return_value
              }
         }
         return return_value
    }
     
     
    
     
     `
    
    return js
}
