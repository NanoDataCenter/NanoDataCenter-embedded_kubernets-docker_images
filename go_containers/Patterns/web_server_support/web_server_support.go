package web_support


//import "fmt"
import (
  "strings"  
  "html/template"
  "net/http"
  //"path/filepath"
)


type Menu_function_type func(http.ResponseWriter, *http.Request)



type Menu_element struct{
  
    Menu_name string;
    Menu_link string;
    Menu_function  Menu_function_type  
    
    
}

type Menu_array []Menu_element





var template_array *template.Template


func Init_web_support()   {
    
   
    generate_base_templates()
    
    
}

func Register_web_pages( menu_array Menu_array){
 
    for _,element := range menu_array {
       
       http.HandleFunc(element.Menu_link,element.Menu_function)
    }
    
    
}

func Generate_single_row_menu( menu_array Menu_array )*template.Template {
    
    working_array := make([]string,0)
    working_array = append(working_array,nav_start)
    for _,menu_element := range menu_array {
        if menu_element.Menu_name != "/"{
            
            element := `<a class="dropdown-item" href="`+menu_element.Menu_link+`"  target="_self">`+menu_element.Menu_name+"</a>"
            working_array = append(working_array,element)
        }
    }
    working_array = append(working_array,nav_end)
    template_string :=strings.Join(working_array,"\n")
    
    template.Must(template_array.New("menu").Parse(template_string)) 
    return template_array
}
    




func generate_base_templates() {
    template_array = template.New("base")
    template.Must(template_array.New("bootstrap").Parse(base_template))
    
}




var base_template = `

<!DOCTYPE html>
<html lang="en">
  <head>
   
    
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/css/bootstrap.min.css" integrity="sha384-TX8t27EcRE3e/ihU7zmQxVncDAy5uIKz4rEkgIXeMed4M0jlfIDPvg6uqKI2xXr2" crossorigin="anonymous">
<style>
input[type="range"] {
  width: 100%;
  margin-bottom: 3rem;
}

.auto {
  -ms-touch-action: auto;
  touch-action: auto;
} 
</style>
<title>{{.Title}}</title>
  </head>

  <body>
  {{template "menu" }}
  
   
  {{template "application"}}
  
 
   

    <!-- jquery & Bootstrap JS -->
<script src="https://code.jquery.com/jquery-3.5.1.slim.min.js" integrity="sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj" crossorigin="anonymous"></script>
<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-ho+j7jyWK8fNQe+A12Hb8AhRq26LrZ/JpcUGGOn+Y7RsweNrtN/tE3MoK7ZeZDyx" crossorigin="anonymous"></script>

<script type="text/javascript" >

</script>
  </body>
</html>

`

var nav_start = `
<nav class="navbar navbar-expand-sm bg-dark navbar-dark">

  <!-- Links -->
  <ul class="navbar-nav">

    <!-- Dropdown -->
    <li class="nav-item dropdown">
      <a class="nav-link dropdown-toggle" href="#" id="navbardrop" data-toggle="dropdown">Menu</a>
      
      <div class="dropdown-menu">
`



        
var nav_end = `
      </div>
    </li>
  </ul>
  <ul class="navbar-nav">

      <button id="status_panel", class="btn " type="submit">Status</button>
  </ul>
  <nav class="navbar navbar-light bg-dark navbar-dark">
  <span class="navbar-text" >
   <h4 id ="status_display"> Status: </h4>
  </span>
  </nav>
</nav>

`
