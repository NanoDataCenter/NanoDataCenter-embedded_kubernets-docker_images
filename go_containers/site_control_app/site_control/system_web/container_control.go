package site_web_server

import (
   
    "net/http"
    "html/template"
    //"lacima.com/Patterns/web_server_support/jquery_react_support"
)




var container_control_template *template.Template

var container_control_html = `
<div class="container">
  <div class="jumbotron">
    <h1>Welcome to my site!</h1>
    <p>This is where I would normally ask you to sign up for something.</p>
  </div>
</div>
`


func container_control_init(){
    container_control_template ,_ = base_templates.Clone()
    
    template.Must(container_control_template.New("application").Parse(container_control_html))
    
    
}    



func container_control(w http.ResponseWriter, r *http.Request) {
   data := make(map[string]interface{})
   data["Title"] = "Container Control"
   container_control_template.ExecuteTemplate(w,"bootstrap", data)
   
    
    
   
   
}

/*

window.location.href = url;
location.reload();

$( "#target" ).click(function() {
  alert( "Handler for .click() called." );
});

ckb = $("#ickb").is(':checked');

<div class="form-check form-check-inline">
      <input id="checkbox3" type="checkbox" checked="checked">
      <label for="checkbox3">Checkbox checked</label>
    </div>

<div class="form-check form-check-inline">
  <input class="form-check-input" type="checkbox" id="inlineCheckbox1" value="option1">
  <label class="form-check-label" for="inlineCheckbox1">1</label>
</div>
<div class="form-check form-check-inline">
  <input class="form-check-input" type="checkbox" id="inlineCheckbox2" value="option2">
  <label class="form-check-label" for="inlineCheckbox2">2</label>
</div>
<div class="form-check form-check-inline">
  <input class="form-check-input" type="checkbox" id="inlineCheckbox3" value="option3" disabled>
  <label class="form-check-label" for="inlineCheckbox3">3 (disabled)</label>
</div>
*/
