
package site_web_server

import (
    
    "net/http"
    "html/template"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
)


var base_templates *template.Template

func Init_site_web_server(){
   
   init_web_server_pages()
   go http.ListenAndServe(":80", nil)
}



func init_web_server_pages() {
    web_support.Init_web_support(introduction_page)
    base_templates = define_web_pages()
    initialize_handlers()
   
}







func define_web_pages()*template.Template  {
 
    return_value := make(web_support.Menu_array,0)
    

    menu_element := web_support.Menu_element{ "introduction page","/introduction_page",introduction_page}
    return_value = append(return_value,menu_element)   

    
    menu_element = web_support.Menu_element{ "application server","/application_server", application_servers}
    return_value = append(return_value,menu_element)

    menu_element = web_support.Menu_element{ "node status","/node_status", node_status}
    return_value = append(return_value,menu_element)    
    
    menu_element = web_support.Menu_element{ "container status","/container_status",container_status}
    return_value = append(return_value,menu_element)        
    
    menu_element = web_support.Menu_element{ "container control","/container_control",container_control}
    return_value = append(return_value,menu_element)    
 
    menu_element = web_support.Menu_element{ "system control","/system_control",system_control}
    return_value = append(return_value,menu_element)    
    
    web_support.Register_web_pages(return_value)
    return web_support.Generate_single_row_menu(return_value)
    
}












func initialize_handlers(){
 
    introduction_page_init()
    application_servers_init()
    node_status_init()
    container_status_init()
    container_control_init()
    system_control_init()
    
    
}

var introduction_page_template *template.Template

var introduction_page_html = `
<div class="container">
  <div class="jumbotron">
    <h1>Welcome to my site!</h1>
    <p>This is where I would normally ask you to sign up for something.</p>
  </div>
</div>
`


func introduction_page_init(){
    introduction_page_template ,_ = base_templates.Clone()
    
    template.Must(introduction_page_template.New("application").Parse(introduction_page_html))
    
    
}    



func introduction_page(w http.ResponseWriter, r *http.Request) {
   data := make(map[string]interface{})
   data["Title"] = "Introduction"
   introduction_page_template.ExecuteTemplate(w,"bootstrap", data)
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
}
*/


 



