package rain_history_page


import (
   "fmt"
   "time"
   "strings"
   "sort"
   "encoding/json"
  //"html/template"
  //"net"
  "net/http"
  "lacima.com/server_libraries/postgres" 
  //"lacima.com/go_application_containers/irrigation/eto/eto_calc/eto_support"
  "html/template"
   "github.com/vmihailenco/msgpack/v5"
   "lacima.com/Patterns/web_server_support/jquery_react_support"
)

var base_templates  *template.Template
var  rain_history     pg_drv.Postgres_Stream_Driver

func Init(templates *template.Template, history pg_drv.Postgres_Stream_Driver){
    base_templates = templates
    rain_history    = history
    
}


func Generate_page_table(w http.ResponseWriter, r *http.Request){

    title := "<center><H3>Display of Rain Values in Table Form</H3></center><BR>"
    html := title+generate_html_text()
    working_template,_ := base_templates.Clone()
    template.Must(working_template.New("application").Parse(html))
    data := make(map[string]interface{})
    data["Title"] = "Rain HISTORICAL DATA In TABLE FORM"
    working_template.ExecuteTemplate(w,"bootstrap", data)    
    
}

func Generate_page_graph(w http.ResponseWriter, r *http.Request){

    title := "<center><H3>Display of Rain Values in Graph Form</H3></center><BR>"
    html := title+generate_html_graph()
    working_template,_ := base_templates.Clone()
    template.Must(working_template.New("application").Parse(html))
    data := make(map[string]interface{})
    data["Title"] = "Rain HISTORICAL DATA In TABLE FORM"
    working_template.ExecuteTemplate(w,"bootstrap", data)    
    
}


func generate_html_text()string{
    keys,_ := rain_history.Find_Tag1_Entries() 
    sort.Strings(keys)
    return_values := make([]string,len(keys))
    for index,key := range keys{
       return_values[index] = generate_text_table(key)
    }
    return_value := strings.Join(return_values,"\n")
    return return_value
}


func generate_text_table(key string )string {
 
     ten_year     := int64(3600*24*365*10)
     db_data, status := rain_history.Select_after_time_stamp_desc_tag1(key, ten_year)
     if status == true {
         title := "<BR><BR><h4>Station  "+key+ "   Historical Data</h4>"
         display_list := make([][]string,len(db_data))
         for index,data := range db_data{
            t :=  time.Unix(data.Time_stamp/1000000000,0)
            string_date := t.Format(time.UnixDate) 
            data ,flag := decode_data(data.Data)
            if flag == true {
               display_list[index] = []string{key,string_date,data} 
            }
         }
         return title+web_support.Setup_data_table("rain_history_"+key,[]string{"Station","Date","Value"},display_list)
     }
     return "<BR><BR><h4>Station "+key+ "  Has No Data</h4>"
    
}

func generate_html_graph()string{
    keys,_ := rain_history.Find_Tag1_Entries()
    sort.Strings(keys)
    return_values := make([]string,len(keys))
    for index,key := range keys{
       return_values[index] = generate_data_graph(index,key)
    }
    return_value := strings.Join(return_values,"\n")
    return return_value
}


func generate_data_graph(index int, key string )string {
 
     ten_year     := int64(3600*24*365*10)
     db_data, status := rain_history.Select_after_time_stamp_desc_tag1(key, ten_year)
     
     if status == true {
         title := "<BR><BR><h4>Station  "+key+ "   Historical Data</h4>"+`<script src="https://cdn.plot.ly/plotly-2.4.2.min.js"></script>`
         time_trace       := make([]int64,len(db_data))
         data_trace     := make([]float64,len(db_data))
         for index,data := range db_data{
            time_trace[index]            = data.Time_stamp
            data_trace[index]            = decode_data_float(data.Data)
            
         }
         temp_time,_                     :=  json.Marshal(time_trace)
         json_time_trace                 :=  string(temp_time)
         temp_data,_                     :=  json.Marshal(data_trace)
         json_data_trace                 :=  string(temp_data)
         return title+generate_graph_data(index,json_time_trace, json_data_trace)
     }
     return "<BR><BR><h4>Station "+key+ "  Has No Data</h4>"
    
}


func generate_graph_data(index int, json_time_trace,json_data_trace string )string{
    
    index_string := fmt.Sprintf("_%d",index)
    
    return_value := `<div id="z_panel`+index_string + `" style="width:100%;height:300pt;"></div>`
   
    return_value = return_value + `
    <script>
    `+
    `json_time_trace="`+ string(json_time_trace)+`";`+`
    
     json_data_trace="`+ string(json_data_trace)+`";`+`
     `+

    
    `   
    var time_trace         =  JSON.parse(json_time_trace)
    var data_trace            =  JSON.parse(json_data_trace)
    

    
    for( i = 0;i< time_trace.length; i++)
    { time_trace[i] = new Date(time_trace[i]/1e6)
    } 
    



     
     
    var trace1 = {
        x: time_trace,
        y: data_trace,
        name: 'Rain Data',
        mode: 'lines+markers',
       type: 'scatter'
     };
    

     
    var data = [trace1];

    
    
    
	z_panel = document.getElementById('z_panel`+index_string+`');
	Plotly.newPlot( z_panel,data );
    </script>  `
    
    return return_value
    
}

func decode_data( msgpack_data string)(string ,bool){
    
    item := make(map[string]interface{})
    err := msgpack.Unmarshal([]byte(msgpack_data), &item)
    if err != nil {
        panic("bad msgpack")
    }
    value := item["value"].(float64)
    if value < .001 {
        return "",false
    }
    value_string := fmt.Sprintf("%.2f",value)
    return value_string, true

}

func decode_data_float( msgpack_data string)float64{
    
    item := make(map[string]interface{})
    err := msgpack.Unmarshal([]byte(msgpack_data), &item)
    if err != nil {
        panic("bad msgpack")
    }
    value := item["value"].(float64)
    
    return value

}
