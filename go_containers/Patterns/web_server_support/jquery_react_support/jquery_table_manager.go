package web_support

import (
    "strings"
)


func Table_Generate_table(title,id_tag string)string{
    
    return_value := make([]string,2)
    return_value[0] =`<h4>`+title+`</h4>`
    return_value[1] = `<table id="`+id_tag+`" class="display" width="100%"></table>`
    return strings.Join(return_value,"")
}


func Table_JS_Routines()string{
    js := `
    <script type="text/javascript"> 
    var Table_column_functions = {}
    var Table_data_build               = {}
    function Table_create_table( tag,columns,column_functions){
     if( columns.length != column_functions.length){
         alert("column definitions not equal "+tag)
         return
      }
      Table_column_functions[tag] = column_functions
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
  
  function Table_clear_table(tag){
      Table_data_build[tag] = []
   }
  
  function Table_add_row(tag,data){
         let temp  = []
         let column_functions = Table_column_functions[tag]
       
         for( i= 0; i<data.length;i++){
              temp.push( column_functions[i](data[i] ))
             
         }
        Table_data_build[tag].push(temp)
    }
  
  
  function Table_load_table(tag){
    
     let table = $(tag).DataTable()
     table.clear()
     table.rows.add(  Table_data_build[tag] )
     table.draw()
  
  
 }
    
 function Table_to_string( value ){
   return String(value)
}
 
 function Table_to_json( value){
      return JSON.stringify(value)
  }
  
  function Table_check_box_element(key){
  
        let label = " "
        check_box = '<div class="form-check">\n'+
       '<label class=""btn  btn-toggle" " for="'+key+'">\n'+     	   
       '<input type="checkbox" class="btn  btn-toggle" id="'+key+'" name="optradio" value='+key+'>'+label+ '</label></div>'
        return check_box
  }
  
  function Table_radio_button_element(key){
  
       let label = " "
       radio_button = '<div class="form-check">\n'+
       '<label class=""btn  btn-toggle" " for="'+key+'">\n'+     	   

       '<input type="radio" class="btn  btn-toggle"  id="'+key+'" name="optradio" value='+key+'>'+label+ '</label></div>'
       return radio_button
   }
  
   function Table_find_select_index(base_id,length){
       
       for( let i=0; i< length; i++)
       {
         if( $("#"+base_id+i).is(':checked') == true )
         {
           return i;
         }
      }
      return  -1;  // no item selected
     }
     
     function Table_find_check_box_elements(base_id,length){
       var return_value = []
       for( let i=0; i< length; i++)
       {
         if( $("#"+base_id+i).is(':checked') == true )
         {
           return_value.push(i)
         }
      }
      return return_value
     
     
     }
    
        function Table_find_check_box_elements(base_id,length){
       var return_value = []
       for( let i=0; i< length; i++)
       {
         if( $("#"+base_id+i).is(':checked') == true )
         {
           return_value.push(i)
         }
      }
      return return_value
     
     
     }
     
    function Table_unselect_check_box_elements(base_id,length){
      
       for( let i=0; i< length; i++)
       {
         $("#"+base_id+i).prop('checked', false)
       
      }
      
     
     
     }
    function Table_select_check_box_elements(base_id,length){
      
       for( let i=0; i< length; i++)
       {
         $("#"+base_id+i).prop('checked', true)
       
      }
     
     
     
     }
    
    function Table_construct_table_elements(table,indexes){
        return_value = []
        for(let i=0; i<indexes.length;i++){
               let index = indexes[i]
               return_value.push(table[index])
         }
        return return_value
     }
    
    
    function Table_do_move(length,base_id,move_id){
   
    
    let move_position = Table_find_select_index(base_id,length)
    if( move_position == -1 ){
        alert("Move Position not chosen")
        return
    }
    let move_steps = Table_find_check_box_elements(move_id,length)
    if( move_steps.length == 0 ){
        alert("Move Steps not chosen")
        return
     }
     
     return  Table_calculate_move(length,move_position,move_steps)
}
    
    
    function Table_calculate_move(length,move_position,move_array){
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
             return_valve = Table_inner_move(temp,return_value)
             
         }
        
        return return_value
     }
     
     function Table_inner_move(value,return_value){
        for(let j=0;j<return_value.length;j++){
              if( return_value[j] == -1 ){
                return_value[j] =value
                return return_value
              }
         }
         return return_value
    }
    
    function Table_remap_table( table_data, new_indexes){
         return_value = []
         for(let i=0;i<new_indexes.length;i++){
             let new_element = table_data[new_indexes[i]]
             return_value.push(new_element)
          }
         return return_value
}
     </script>
    
     
     `
    
    return js
}
