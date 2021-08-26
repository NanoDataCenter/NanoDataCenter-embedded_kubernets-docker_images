package web_support

import "strings"
import "strconv"
//import "fmt"



const alink_start  string  = "<a href="


const  alink_end string    = "</a>"


type Link_type struct{
 
    Display string;
    Link    string;
    
}

type Accordion_Elements struct{

    Title string;
    Body  string;
    
}



//  <a href="#" class="list-group-item list-group-item-action">A fourth link item</a>
/*
 The link field must have the following form
 "\"http://127.0.0.1/\"  target=\"_blank\""
 */  
func Generate_list_link_component( container_id string, header string, link_array []Link_type )string{
    html_statements := make([]string,0)
    html_statements = append(html_statements,"<div id=\""+container_id+ "\" class=\"container\">")
    html_statements = append(html_statements,"<h2>"+header+"</h2>")
    html_statements = append(html_statements,"<div class=\"list-group\">")
    
    
    for _,element := range link_array {
       html_statements = append(html_statements,alink_start+element.Link+"  class=\"list-group-item list-group-item-action\">"+element.Display+alink_end)     
    }    
    
    html_statements = append(html_statements,"</div>")
    html_statements = append(html_statements,"</div>")
    return strings.Join(html_statements,"\n")
    


    
}  





// <li class="list-group-item">And a fifth one</li>
func Generate_list_link( container_id string, header string, display_array []string )string{
    html_statements := make([]string,0)
    html_statements = append(html_statements,"<div id=\""+container_id+ "\" class=\"container\">")
    html_statements = append(html_statements,"<h2>"+header+"</h2>")
    html_statements = append(html_statements,"<ul class=\"list-group\">")
    
    
    for _,element := range display_array {
       html_statements = append(html_statements,"<li class=\"list-group-item\">"+element+"</li>")     
    }    
    
    html_statements = append(html_statements,"</ul>")
    html_statements = append(html_statements,"</div>")
     return strings.Join(html_statements,"\n")
    


    
} 


func Populate_accordian_elements( title_array,body_array []string )[]Accordion_Elements{
    
      return_value := make([]Accordion_Elements,0)
    
      for index, title := range title_array {
          return_value = append(return_value, Accordion_Elements{ Title:title, Body:body_array[index] })
      }
    
      return return_value   
}    

func Generate_accordian(container_id , title string, data_elements []Accordion_Elements )string{
    
    html_statements := make([]string,0)
    html_statements = append(html_statements,"<center> <h2>"+title+"</h2> </center>")
    html_statements = append(html_statements,`<div class="accordion" id="`+container_id+`">`)
    for count,element := range data_elements {
        header_id := container_id+"header"+strconv.Itoa(count)
        collaspe_id := container_id+"collaspe"+strconv.Itoa(count)
  
        html_statements = append(html_statements,`<div class="card">`)
        html_statements = append(html_statements,`<div class="card-header" id="`+header_id+`">`)
        html_statements = append(html_statements,`<h2 class="mb-0">`)
        text1 := `<button class="btn btn-link btn-block text-left" type="button" data-toggle="collapse" data-target="#`+collaspe_id
        
        if count == 0 {
             
             text2 := `" class="collapse show" aria-controls="`+collaspe_id+`">`
             html_statements = append(html_statements,text1+text2)
        }else{    
           
           text2 := `" class="collapse " aria-controls="`+collaspe_id+`">`
           html_statements = append(html_statements,text1+text2)
        }
        
        html_statements = append(html_statements,element.Title)
        html_statements = append(html_statements,`</button>`)
        html_statements = append(html_statements,`</h2>`)
        html_statements = append(html_statements,`</div>`)
              
       
 


        text := `<div id="`+collaspe_id+`" class="collapse " aria-labelledby="`+header_id+`" data-parent="#`+container_id+`">`
        html_statements = append(html_statements,text)
        html_statements = append(html_statements,`<div class="card-body ">`)
        html_statements = append(html_statements,element.Body)
        html_statements = append(html_statements,`</div>`)
        html_statements = append(html_statements,`</div>`)
        html_statements = append(html_statements,`</div>`)    
        
    }
    html_statements = append(html_statements,"</div>")
    
     return strings.Join(html_statements,"\n")
}



/*
 <div class="accordion" id="accordionExample">
  <div class="card">
    <div class="card-header" id="headingOne">
      <h2 class="mb-0">
        <button class="btn btn-link btn-block text-left" type="button" data-toggle="collapse" data-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne">
          Collapsible Group Item #1
        </button>
      </h2>
    </div>

    <div id="collapseOne" class="collapse show" aria-labelledby="headingOne" data-parent="#accordionExample">
      <div class="card-body">
        Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry richardson ad squid. 3 wolf moon officia aute, non cupidatat skateboard dolor brunch. Food truck quinoa nesciunt laborum eiusmod. Brunch 3 wolf moon tempor, sunt aliqua put a bird on it squid single-origin coffee nulla assumenda shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore wes anderson cred nesciunt sapiente ea proident. Ad vegan excepteur butcher vice lomo. Leggings occaecat craft beer farm-to-table, raw denim aesthetic synth nesciunt you probably haven't heard of them accusamus labore sustainable VHS.
      </div>
    </div>
  </div>
  <div class="card">
    <div class="card-header" id="headingTwo">
      <h2 class="mb-0">
        <button class="btn btn-link btn-block text-left collapsed" type="button" data-toggle="collapse" data-target="#collapseTwo" aria-expanded="false" aria-controls="collapseTwo">
          Collapsible Group Item #2
        </button>
      </h2>
    </div>
    <div id="collapseTwo" class="collapse" aria-labelledby="headingTwo" data-parent="#accordionExample">
      <div class="card-body">
        Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry richardson ad squid. 3 wolf moon officia aute, non cupidatat skateboard dolor brunch. Food truck quinoa nesciunt laborum eiusmod. Brunch 3 wolf moon tempor, sunt aliqua put a bird on it squid single-origin coffee nulla assumenda shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore wes anderson cred nesciunt sapiente ea proident. Ad vegan excepteur butcher vice lomo. Leggings occaecat craft beer farm-to-table, raw denim aesthetic synth nesciunt you probably haven't heard of them accusamus labore sustainable VHS.
      </div>
    </div>
  </div>
  <div class="card">
    <div class="card-header" id="headingThree">
      <h2 class="mb-0">
        <button class="btn btn-link btn-block text-left collapsed" type="button" data-toggle="collapse" data-target="#collapseThree" aria-expanded="false" aria-controls="collapseThree">
          Collapsible Group Item #3
        </button>
      </h2>
    </div>
    <div id="collapseThree" class="collapse" aria-labelledby="headingThree" data-parent="#accordionExample">
      <div class="card-body">
        Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry richardson ad squid. 3 wolf moon officia aute, non cupidatat skateboard dolor brunch. Food truck quinoa nesciunt laborum eiusmod. Brunch 3 wolf moon tempor, sunt aliqua put a bird on it squid single-origin coffee nulla assumenda shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore wes anderson cred nesciunt sapiente ea proident. Ad vegan excepteur butcher vice lomo. Leggings occaecat craft beer farm-to-table, raw denim aesthetic synth nesciunt you probably haven't heard of them accusamus labore sustainable VHS.
      </div>
    </div>
  </div>
</div>
*/
