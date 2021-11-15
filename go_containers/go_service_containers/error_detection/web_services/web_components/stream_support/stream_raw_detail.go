
package stream_support

import (
    "fmt"
    "strings"
    //"strconv"
    //"sort"
    //"time"
    "net/http"
    "encoding/json"
  
    "html/template"

   "lacima.com/Patterns/web_server_support/jquery_react_support"
   "lacima.com/Patterns/msgpack_2"
   "github.com/montanaflynn/stats"
 
 


)

 



type raw_stream_type struct{
    value       float64  
    time_stamp  int64
}

var return_value  []raw_stream_type
var data_x   []int64
var data_y   []float64

var stats_input  []stats.Coordinate
var linear_x  []float64
var linear_y  []float64

func stream_raw_detail  (w http.ResponseWriter, r *http.Request) {
   
   key := r.URL.Query().Get("key")
   
   stream_raw_template ,_ := base_templates.Clone()
   
   generate_data(key)
   display_data := stream_raw_generate_html()

 
    
   data := make(map[string]interface{})
   key_list      := strings.Split(key,"~+~")
   key_display   := strings.Join(key_list,"~")
   data["Title"] = key_display
   
   
   header_display := "<center><h2>Detail Data For </h2></center><br>"
   header_display =header_display+ "<center><h3>"+ key_display+ "</h3></center><br>"
   template.Must(stream_raw_template.New("application").Parse(header_display+display_data))
   
   
   stream_raw_template.ExecuteTemplate(w,"bootstrap",data )
   
}



func stream_raw_generate_html()string{ 

    
    return_value := `
<script src="https://cdn.plot.ly/plotly-2.4.2.min.js"></script>
<p>
 <button class="btn btn-primary" id="btn1" type="button" >
    Graphic Display
  </button>
  <button class="btn btn-primary" id="btn2" type="button" >
    Table Display
  </button>
</p>
<div  id="graphic">
  <div class="card card-body">
  `+ generate_graphic_display( )+`
  
  </div>
</div>
<div  id="tabular">
  <div class="card card-body">
   `+ generate_table_display() +`
   
  </div>
</div>

<script>

$("#graphic").show();
$("#tabular").hide();

$("#btn1").on("click", function(){
   
    $("#graphic").show();
    $("#tabular").hide();
});
$("#btn2").on("click", function(){
   
    $("#tabular").show();
    $("#graphic").hide();
});
</script>
`
    return return_value

}

func generate_data(key string){
    
    key_tags := strings.Split(key,"~+~")
    
    if len(key_tags) <   5 {
       panic("bad postgres key")
    }
    
    tag1 := key_tags[0]
    tag2 := key_tags[1]
    tag3 := key_tags[2]
    tag4 := key_tags[3]
    tag5 := key_tags[4]
    
    
    where_clause   := " tag1 = '"+tag1+"' and tag2 = '"+tag2+"' and tag3 = '"+tag3+"' and tag4 = '"+tag4+"' and tag5 = '"+tag5+"'  and  time >= 0 ORDER BY time DESC LIMIT 200 "
    pg_data,status := stream_control.process_data_stream.Select_where(where_clause)
 
    return_value = make([]raw_stream_type,len(pg_data))
    if status != true {
        panic("bad select")
    }
    for index,data_element := range pg_data{
       
       value,err      :=    msg_pack_utils.Unpack_float64(data_element.Data)
       if err != true {
           panic("bad packed data")
       }
    
       return_value[index].value = value
       return_value[index].time_stamp = data_element.Time_stamp
    }
    data_x = make([]int64,len(return_value))
    data_y = make([]float64,len(return_value))
    stats_input = make([]stats.Coordinate,len(return_value))
    linear_x = make([]float64,len(return_value))
    linear_y = make([]float64,len(return_value))
    for index,value := range return_value{
        data_x[index]      = value.time_stamp
        data_y[index]      = value.value
        stats_input[index] = stats.Coordinate{float64(value.time_stamp),value.value}
    }
    lr_output,_  :=  stats.LinearRegression(stats_input)
    for index , value := range lr_output{
        linear_x[index] = value.X
        linear_y[index] = value.Y
    }
}






func generate_table_display( )string {
    
    display_list := make([][]string,len(return_value))  
       
    for index,data            := range return_value {
        value                 :=   fmt.Sprintf("%f",data.value)
        time_stamp            :=   format_time(data.time_stamp, true) 
        display_list[index] = []string{value,time_stamp} 
    
    }
    return web_support.Setup_data_table("raw_list",[]string{"VALUE","TIME" },display_list)
    
}
        
        

func generate_graphic_display( ) string {
    
    json_data_x,_ :=  json.Marshal(data_x)
    json_data_y,_ :=  json.Marshal(data_y)
    json_linear_x,_ :=  json.Marshal(linear_x)
    json_linear_y,_ :=  json.Marshal(linear_y)    
    
    
    
    
    return_value := `<div id="graphic_panel" style="width:100%;height:300pt;"></div>`
   
    return_value = return_value + `
    <script>
    `+
    `json_data_x="`+ string(json_data_x)+`";`+`
     `+
    `json_data_y="`+ string(json_data_y)+`"; `+
    
    `
     `+
    `json_linear_x="`+ string(json_linear_x)+`";`+`
     `+
    `json_linear_y="`+ string(json_linear_y)+`"; `+
    
    `   
    var trace1_x  =  JSON.parse(json_data_x)
    var trace1_y  =  JSON.parse(json_data_y) 
    
    var trace2_x  =  JSON.parse(json_linear_x)
    var trace2_y  =  JSON.parse(json_linear_y) 
    
    for( i = 0;i< trace1_x.length; i++)
    { 
        trace1_x[i] = new Date(trace1_x[i]/1e6)
    } 
    
        for( i = 0;i< trace2_x.length; i++)
    { 
        trace2_x[i] = new Date(trace2_x[i]/1e6)
    } 

    var trace1 = {
        x: trace1_x,
        y: trace1_y,
        name: 'Data',
        mode: 'markers',
       type: 'scatter'
     };
    var trace2 = {
        x: trace2_x,
        y: trace2_y,
        name: 'LR',
        mode: 'lines+markers',
       type: 'scatter'
     };
    
    var data = [trace1,trace2];

    
    
    
	graphic_panel = document.getElementById('graphic_panel');
	Plotly.newPlot( graphic_panel,data );
    </script>  `
    
    return return_value
    
}

