package web_support

import "strings"





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
       html_statements = append(html_statements,alink_start+element.Link+"class=\"list-group-item list-group-item-action\">"+element.Display+alink_end)     
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


func Generate_Accordian(container_id string ,  data_elements []Accordion_Elements )string{
    
    html_statements := make([]string,0)
    html_statements = append(html_statements,"<div class=\"accordion accordion-flush\" id=\""+container_id+"\">")
    for count,element := range data_elements {
       button_id := container_id+string(count)
       html_statements = append(html_statements,"<div class=\"accordion-item\">")
       html_statements = append(html_statements,"<h2 class=\"accordion-header\" id=\""+container_id+string(count)+"\">")
       temp1 := "<button class=\"accordion-button collapsed\" type=\"button\" data-bs-toggle=\"collapse\" data-bs-target=\"#"
       temp2 := "\" aria-expanded=\"false\" aria-controls=\""+button_id+"\">"
       html_statements = append(html_statements,temp1+button_id+temp2)
       html_statements = append(html_statements,element.Title)
       html_statements = append(html_statements,"</button>")
       html_statements = append(html_statements,"</h2>")
       temp1 = "<div id=\"flush-collapseThree\" class=\"accordion-collapse collapse\" aria-labelledby=\""+button_id+"\""
       temp2 = "\" data-bs-parent=\""+container_id+"\">"
       html_statements = append(html_statements,temp1+temp2)
       html_statements = append(html_statements,element.Body)
       html_statements = append(html_statements,"</div>")
       html_statements = append(html_statements,"</div>")
    }
    html_statements = append(html_statements,"</div>")
     return strings.Join(html_statements,"\n")
}


/*
     <div class="accordion accordion-flush" id="accordionFlushExample">
  <div class="accordion-item">
    <h2 class="accordion-header" id="flush-headingOne">
      <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#flush-collapseOne" aria-expanded="false" aria-controls="flush-collapseOne">
        Accordion Item #1
      </button>
    </h2>
    <div id="flush-collapseOne" class="accordion-collapse collapse" aria-labelledby="flush-headingOne" data-bs-parent="#accordionFlushExample">
      <div class="accordion-body">Placeholder content for this accordion, which is intended to demonstrate the <code>.accordion-flush</code> class. This is the first item's accordion body.</div>
    </div>
  </div>
  <div class="accordion-item">
    <h2 class="accordion-header" id="flush-headingTwo">
      <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#flush-collapseTwo" aria-expanded="false" aria-controls="flush-collapseTwo">
        Accordion Item #2
      </button>
    </h2>
    <div id="flush-collapseTwo" class="accordion-collapse collapse" aria-labelledby="flush-headingTwo" data-bs-parent="#accordionFlushExample">
      <div class="accordion-body">Placeholder content for this accordion, which is intended to demonstrate the <code>.accordion-flush</code> class. This is the second item's accordion body. Let's imagine this being filled with some actual content.</div>
    </div>
  </div>
  <div class="accordion-item">
    <h2 class="accordion-header" id="flush-headingThree">
      <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#flush-collapseThree" aria-expanded="false" aria-controls="flush-collapseThree">
        Accordion Item #3
      </button>
    </h2>
    <div id="flush-collapseThree" class="accordion-collapse collapse" aria-labelledby="flush-headingThree" data-bs-parent="#accordionFlushExample">
      <div class="accordion-body">Placeholder content for this accordion, which is intended to demonstrate the <code>.accordion-flush</code> class. This is the third item's accordion body. Nothing more exciting happening here in terms of content, but just filling up the space to make it look, at least at first glance, a bit more representative of how this would look in a real-world application.</div>
    </div>
  </div>
</div>
*/
