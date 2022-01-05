package eto_daily_variation_page

import (
   "fmt"
   "strings"
   "sort"
   
  //"html/template"
  //"net"
  "net/http"
  
  "lacima.com/go_application_containers/irrigation/eto/eto_calc/eto_support"
  "html/template"
    "lacima.com/redis_support/redis_handlers" 
   "lacima.com/Patterns/web_server_support/jquery_react_support"
)

var base_templates       *template.Template
var  eto_stream_data     redis_handlers.Redis_Hash_Struct

func Init(templates *template.Template, stream_data redis_handlers.Redis_Hash_Struct){
    base_templates = templates
    eto_stream_data    = stream_data
    
}


func Generate_page(w http.ResponseWriter, r *http.Request){
    html := generate_html()
    working_template,_ := base_templates.Clone()
    template.Must(working_template.New("application").Parse(html))
    data := make(map[string]interface{})
    data["Title"] = "ETO Daily Values"
    working_template.ExecuteTemplate(w,"bootstrap", data)    
    
}

/*
func generate_html()string{
    keys := eto_support.Stream_HKeys()    
    sort.Strings(keys)

    display_list := make([][]string,len(keys))
    for index,key := range keys{
       data                          := eto_support.Stream_HGet(key)
       humidity                      := fmt.Sprintf("%.2f",data["humidity"])
       temp_C                        := fmt.Sprintf("%.2f",data["temp_C"])
       SolarRadiationWatts_m_squared := fmt.Sprintf("%.2f",data["SolarRadiationWatts_m_squared"])
       wind_speed                    := fmt.Sprintf("%.2f",data["wind_speed"]) 
       eto_rain_value                := eto_support.ETO_HGet(key)
       eto_value                     := fmt.Sprintf("%.2f",eto_rain_value.Value)
       display_list[index]           = []string{key,humidity,temp_C,SolarRadiationWatts_m_squared,wind_speed,eto_value}
    }
    title := "<CENTER>ETO VALUE COMPARISION</CENTER><BR>"
    return title+web_support.Setup_data_table("eto_parameters",[]string{"Station","HUMIDITY","TEMP C","SOLAR","WIND","ETO"},display_list)
    
}
*/
func generate_html()string{
    
   
    stream_data := make(map[string]map[string][]float64)
    
    html_array := make([]string,5)
    
    keys := eto_support.Stream_HKeys()    
    sort.Strings(keys)
    
    values := make(map[string]float64)
    for _,key := range keys{
       data                               := eto_support.Stream_HGet(key)
       if check_entries( data ) != true {
           break
       }
       stream_data[key] = data
       
       values[key] = eto_support.ETO_HGet(key).Value
       
    }
    
    title := "<CENTER><H3>ETO VALUE COMPARISION</H3></CENTER><BR>"
    html_array[0] = title
    html_array[1] = generate_table(keys,values,"humidity",stream_data)
    html_array[2] = generate_table(keys,values,"temp_C",stream_data)
    html_array[3] = generate_table(keys,values,"SolarRadiationWatts_m_squared",stream_data)
    html_array[4] = generate_table(keys,values,"wind_speed",stream_data)
    return strings.Join(html_array,"<BR>")
    
    
}

func check_entries( input map[string][]float64 ) bool {

    for _, value := range input {
        if len(value) <24 {
            return false
        }
    }
    return true
}
    
    
    
func generate_table( keys []string, values map[string]float64, field string, stream_data map[string]map[string][]float64)string {
    
    title := "<H4>ETO VALUE COMPARISION FOR FIELD "+field+"</H4>"
    display_list := make([][]string,len(keys))
    
    for index,key := range keys {
       display_list[index] = generate_line( key , values[key], field,stream_data)  
        
        
    }
    
    
    return title+web_support.Setup_data_table("eto_parameters_"+field,[]string{"Station","ETO_Value","0","1","2","3","4","5","6","7","8","9","10",
                    "11","12","13","14","15","16","17","18","19","20","21","22","23"},display_list)
    
    
}
 
 
func generate_line( key string, value float64 , field string,  stream_data  map[string]map[string][]float64) []string {
    return_value := make([]string,26)
    return_value[0] = key
    return_value[1] = to_string(value)
    for i := 0; i<24;i++{
       return_value[i+2] = to_string(stream_data[key][field][i])   
        
    }
    return return_value
}    

func to_string( input float64)string{
    
    return    fmt.Sprintf("%.2f",input)
    
}
    

