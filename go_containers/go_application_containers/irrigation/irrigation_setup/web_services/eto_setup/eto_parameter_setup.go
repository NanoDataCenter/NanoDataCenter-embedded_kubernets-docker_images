package eto_setup


import(
    "bytes"
    "text/template"
    //"fmt"
)

type ETO_Parameter_Setup struct{
    
    Save_flag      bool;
    Div_name       string;
    Title          string;
    Field          string;
    Parm_name      string;
    Max_Range      string;
    Step           string;
    Helper1        string;
    Helper2        string;
    Helper3        string;
    
}

var html_parameter_template   string
var js_parameter_template     string



func generate_parameter_setup( input ETO_Parameter_Setup)string{
    generate_html_template()
    generate_js_template()
    html := generate_parameter_setup_html(input) +
            generate_parameter_setup_js(input)
            
    return html
}
    

func generate_parameter_setup_html(input ETO_Parameter_Setup)string{
    t  := template.Must(template.New("Param_html").Parse(html_parameter_template))
   
	var output bytes.Buffer
	err := t.Execute(&output, input)
	if err != nil {
		panic(err)
	}
	return output.String()
    
}

func generate_parameter_setup_js(input ETO_Parameter_Setup)string{
    t := template.Must(template.New("Param_js").Parse(js_parameter_template))
   	var output bytes.Buffer
	err := t.Execute(&output, input)
	if err != nil {
		panic(err)
	}
	return output.String()
    
}





	
func generate_html_template(){
  html_parameter_template  = `




 <div class="container" >
  
   <h2 id="{{.Div_name}}_title">{{.Title}} </h2>

   <div>
        {{if .Save_flag }} 
        <input type="button" id = "{{.Div_name}}_save" value="Save"  data-inline="true"  /> 
        <input type="button" id = "{{.Div_name}}_cancel" value="Cancel" data-inline="true"  />
        {{end}}
        <input type="button" id = "{{.Div_name}}_reset" value="Reset" data-inline="true"  /> 
   </div>
   

   
<div style="margin-top:25px"></div>
<h4 id="{{.Div_name}}_display"></h4>
<div style="margin-top:25px"></div>

<input id = "{{.Div_name}}_input", class="auto" mim = "0" max="{{.Max_Range}}" step="{{.Step}}" type="range">
</div>

`
}

func generate_js_template(){
   js_parameter_template = `    
   <script type="text/javascript">    
   
  var {{.Div_name}}_reset_value
   
  function {{.Div_name}}_load_controls(){
    
    
    {{if .Save_flag }}  
     $("#{{.Div_name}}_save").bind('click',{{.Div_name}}_save)
     $("#{{.Div_name}}_cancel").bind('click',{{.Div_name}}_cancel)
     
     {{end}}
     $("#{{.Div_name}}_reset").bind('click',{{.Div_name}}_reset)
     $("#{{.Div_name}}_input").on('input', {{.Div_name}}_update_display )
  
  }      
  
  function {{.Div_name}}_open( value ){
  
     {{.Div_name}}_reset_value = value
     $("#{{.Div_name}}_input").val( {{.Div_name}}_reset_value)
     {{.Div_name}}_set_slider_title(value)
     
    
  }
  
  function {{.Div_name}}_save() {
  
       var result = confirm("Do you wish to save data?");  
       if( result == true )
       {
           value = parseFloat($("#{{.Div_name}}_input").val())
           update_values("{{.Field}}",value)
           setup_main_screen()
       }

   }
   
   function {{.Div_name}}_cancel()
   {
     setup_main_screen()

   }

   function {{.Div_name}}_set_slider_title( value)
   {
     $("#{{.Div_name}}_display").html( "{{.Helper1}}"+value*{{.Helper2}}+"{{.Helper3}}")
                                
   }

   function {{.Div_name}}_reset()
   {
        $("#{{.Div_name}}_input").val( {{.Div_name}}_reset_value)
        {{.Div_name}}_set_slider_title({{.Div_name}}_reset_value)
      
   }

  function  {{.Div_name}}_get_value(){
  
    value = parseFloat($("#{{.Div_name}}_input").val())
     return ["{{.Field}}",value ]
   }
   
  function {{.Div_name}}_update_display (event) {
    var value
    value = $("#{{.Div_name}}_input").val()
    {{.Div_name}}_set_slider_title( value)
    
   }
</script>
`
}

/*
blank          bool,
    save_flag      bool,
    div_name       string,
    title          string,
    field          string,
    parm_name      string,
    helper_strings [3]string,        

*/
