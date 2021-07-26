
package main

import (
    "fmt"
    "net/http"
"html/template"
  "lacima.com/Patterns/web_server_support"
)


var base_templates *template.Template


func define_web_pages()*template.Template  {
 
    return_value := make(web_support.Menu_array,0)
    
    menu_element := web_support.Menu_element{ "/","/",slash_page}
    return_value = append(return_value,menu_element)
    
    menu_element = web_support.Menu_element{ "hello_intro","/hello_intro",hello_intro}
    return_value = append(return_value,menu_element)   

    menu_element = web_support.Menu_element{ "hello1","/hello1",hello1}
    return_value = append(return_value,menu_element)

    menu_element = web_support.Menu_element{ "hello2","/hello2",hello2}
    return_value = append(return_value,menu_element)    
    
    web_support.Register_web_pages(return_value)
    return web_support.Generate_single_row_menu(return_value)
    
}










func main() {
    web_support.Init_web_support()
    base_templates = define_web_pages()
    initialize_handlers()
   
    http.ListenAndServe(":3000", nil)
}

func initialize_handlers(){
 
    hello_intro_init()
    hello1_init()
    hello2_init()
    
    
}


func slash_page(w http.ResponseWriter, r *http.Request){
    
    if r.URL.Path != "/" {
        errorHandler(w, r, http.StatusNotFound)
        
    }else{
        hello_intro(w,r)
    }
    
}


func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "custom 404")
    }
}


var hello_intro_template *template.Template

var hello_intro_application = `
<div class="container">
  <div class="jumbotron">
    <h1>Welcome to my site!</h1>
    <p>This is where I would normally ask you to sign up for something.</p>
  </div>
</div>
`


func hello_intro_init(){
    hello_intro_template ,_ = base_templates.Clone()
    
    template.Must(hello_intro_template.New("application").Parse(hello_intro_application))
    
    
}    



func hello_intro(w http.ResponseWriter, r *http.Request) {
   data := make(map[string]interface{})
   data["Title"] = "Hello_intro"
   hello_intro_template.ExecuteTemplate(w,"bootstrap", data)
   
    
    
}


var hello1_template *template.Template

var hello1_application = `
<div class="container">
  <div class="jumbotron">
    <h1>hello1</h1>
    <p>+++++++++++++.</p>
  </div>
</div>
`


func hello1_init(){
    hello1_template ,_ = base_templates.Clone()
    
    template.Must(hello1_template.New("application").Parse(hello1_application))
    
    
}    



func hello1(w http.ResponseWriter, r *http.Request) {
   data := make(map[string]interface{})
   data["Title"] = "Hello_intro"
   hello1_template.ExecuteTemplate(w,"bootstrap", data)
   
    
    
}


var hello2_template *template.Template

var hello2_application = `
<div class="container">
  <div class="jumbotron">
    <h1>hello2</h1>
    <p>*****************************************.</p>
  </div>
</div>
`


func hello2_init(){
    hello2_template ,_ = base_templates.Clone()
    
    template.Must(hello2_template.New("application").Parse(hello2_application))
    
    
}    



func hello2(w http.ResponseWriter, r *http.Request) {
   data := make(map[string]interface{})
   data["Title"] = "Hello_intro"
   hello2_template.ExecuteTemplate(w,"bootstrap", data)
   
    
    
}



