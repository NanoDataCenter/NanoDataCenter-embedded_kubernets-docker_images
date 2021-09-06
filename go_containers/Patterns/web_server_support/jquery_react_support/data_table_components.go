package web_support


import "strings"



func Setup_data_table(table_tag string , table_headers []string, table_data[][]string )string{
    html_array    := make([]string,4)
    html_array[0] = generate_table_header(table_tag,table_headers) 
    html_array[1] = generate_table_entries(table_data)
    html_array[2] = generate_table_ending(table_headers)
    html_array[3] = generate_java_script(table_tag)
    return strings.Join(html_array,"\n")

}

func generate_java_script( tag string)string{
    
    return table_script_start + tag + table_script_end 

    
}

const table_script_start string = `
<script>
$(document).ready( function () {
    $('#`
    
const table_script_end string = `').DataTable( { "order": [] , "pageLength": 50 }  );
} );
</script>`
 
func generate_table_header( table_tag string,table_headers []string) string{
  
   return_value := make([]string,7)
   return_value[0] =  `<table id="`+table_tag +`" class="display" style="width:100%">`
   return_value[1] =  "<thead>"
   return_value[2] =  "<tr>"
   return_value[3] =  generate_headers(table_headers)
   return_value[4] = "</tr>"
   return_value[5] = "</thead>"
   return_value[6] = "<tbody>"
   return strings.Join(return_value,"\n")
}


func generate_table_ending( table_headers [] string )string {

   return_value := make([]string,7)
   return_value[0] = "</tbody>"
   return_value[1] = "<tfoot>"
   return_value[2] = "<tr>"
   return_value[3] = generate_headers(table_headers)
   return_value[4] = "</tr>"
   return_value[5] = "</tfoot>"
   return_value[6] = "</table>"
   return strings.Join(return_value,"\n")
}

func generate_headers( table_headers []string) string {
    return_value := make([]string,len(table_headers))
    for index,element := range table_headers {
        return_value[index] = "<th>"+element+"</th>"
    }
    return strings.Join(return_value,"\n")
}
 
       
    
    

 
func generate_table_entries( data_entries [][]string ) string{
    
    return_value := make([]string,len(data_entries))
    
    for index, element := range data_entries {
        return_value[index] = generate_table_element(element)
    }
    return strings.Join(return_value,"\n")
}


func generate_table_element( row_element []string )string{
    return_value := make([]string,len(row_element)+2)
    return_value[0] = "<tr>"
    for index , column_element := range row_element {
        return_value[index +1 ] = "<td>"+column_element+"</td>"
    }
    return_value[len(row_element)+1] = "</tr>"
    return strings.Join(return_value,"\n")
}
    
 
 
 
 
