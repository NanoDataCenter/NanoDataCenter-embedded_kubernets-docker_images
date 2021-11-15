package stream_support


import (
    "fmt"
    "strings"
    "encoding/json"
    //"strconv"
    //"sort"
    //"time"
    "net/http"
     //"net/url"
    "html/template"
    "github.com/vmihailenco/msgpack/v5"
     "lacima.com/Patterns/web_server_support/jquery_react_support"
    //"encoding/json"
    //"lacima.com/redis_support/graph_query"
    //"lacima.com/redis_support/generate_handlers"
   // "lacima.com/redis_support/redis_handlers"
   // "lacima.com/server_libraries/postgres"
   //  "lacima.com/Patterns/msgpack_2"
   //"lacima.com/Patterns/web_server_support/jquery_react_support"
  
    
    
 


)

type filtered_value_type struct{
    filtered_value    float64
    current_value     float64
    current_velocity  float64
    z_value           float64
    std               float64
    z_state           bool  
}    

type filtered_return_value_type struct {
 value        filtered_value_type    
 time_stamp   int64   
}

var    filtered_return_value  []filtered_return_value_type
var    json_time_trace    string
var    json_filter_trace  string
var    json_std_trace     string
var    json_raw_trace     string
var    json_z_trace       string
var    json_velocity_trace      string
var    json_zstate_trace    string

func stream_filtered_detail  (w http.ResponseWriter, r *http.Request) {
   
   key := r.URL.Query().Get("key")
   
   stream_raw_template ,_ := base_templates.Clone()
   
   generate_func_data(key)
   display_data := stream_detail_generate_html()

 
    
   data := make(map[string]interface{})
   key_list      := strings.Split(key,"~+~")
   key_display   := strings.Join(key_list,"~")
   data["Title"] = key_display
   
   
   header_display := "<center><h2>Filtered Data For </h2></center><br>"
   header_display =header_display+ "<center><h3>"+ key_display+ "</h3></center><br>"
   template.Must(stream_raw_template.New("application").Parse(header_display+display_data))
   
   
   stream_raw_template.ExecuteTemplate(w,"bootstrap",data )
   
}

func stream_detail_generate_html()string{
       return_value := `
<script src="https://cdn.plot.ly/plotly-2.4.2.min.js"></script>
<p>
 <button class="btn btn-primary" id="btn1" type="button" >
    Filtered Data
  </button>
  <button class="btn btn-primary" id="btn2" type="button" >
   Z Value
  </button>
   <button class="btn btn-primary" id="btn3" type="button" >
    Velocity Data
  </button>
     <button class="btn btn-primary" id="btn4" type="button" >
    Table Display
  </button>
</p>
<div  id="filtered_data">
  <div class="card card-body">
  `+ generate_filted_data( )+`
  
  </div>
</div>
<div  id="z_display">
  <div class="card card-body">
  `+ generate_z_data( )+`
  
  </div>
</div>
<div  id="v_display">
  <div class="card card-body">
  `+ generate_v_data( )+`
  
  </div>
</div>
<div  id="tabular">
  <div class="card card-body">
   `+ generate_filtered_table_display() +`
   
  </div>
</div>

<script>

$("#iltered_data").show();
$("#z_display").hide();
$("#v_display").hide();
$("#tabular").hide();


$("#btn1").on("click", function(){
   
    $("#filtered_data").show();
    $("#z_display").hide();
    $("#v_display").hide();
    $("#tabular").hide();
});


$("#btn2").on("click", function(){
    $("#filtered_data").hide();
    $("#z_display").show();
    $("#v_display").hide();
    $("#tabular").hide();
});
$("#btn3").on("click", function(){
    $("#filtered_data").hide();
    $("#z_display").hide();
    $("#v_display").show();
    $("#tabular").hide();
});
$("#btn4").on("click", function(){
    $("#filtered_data").hide();
    $("#z_display").hide();
    $("#v_display").hide();
    $("#tabular").show();
});
</script>
`
    return return_value
}


func generate_filted_data( ) string {
    
 
    
    
    
    return_value := `<div id="filtered_panel" style="width:100%;height:300pt;"></div>`
   
    return_value = return_value + `
    <script>
    `+
    `json_time_trace="`+ string(json_time_trace)+`";`+`
     `+
    `json_filter_trace="`+ string(json_filter_trace)+`"; `+
    
    `
     `+
    `json_std_trace="`+ string(json_std_trace)+`";`+`
     `+
    `json_raw_trace="`+ string(json_raw_trace)+`"; `+
    
    `   
    var time_trace  =  JSON.parse(json_time_trace)
    var filter_trace  =  JSON.parse(json_filter_trace) 
    var std_trace     =  JSON.parse(json_std_trace)
    var raw_trace     =  JSON.parse(json_raw_trace)

    
    for( i = 0;i< time_trace.length; i++)
    { time_trace[i] = new Date(time_trace[i]/1e6)
    } 
    


    var trace1 = {
        x: time_trace,
        y: filter_trace,
        error_y: {
          type: 'data',
          array: std_trace,
          visible: true
         },
        name: 'FILTERED',
        mode: 'lines+markers',
       type: 'scatter'
     };
     
     
    var trace2 = {
        x: time_trace,
        y: raw_trace,
        name: 'RAW',
        mode: 'lines+markers',
       type: 'scatter'
     };
    
    var data = [trace1,trace2];

    
    
    
	graphic_panel = document.getElementById('graphic_panel');
	Plotly.newPlot( filtered_panel,data );
    </script>  `
    
    return return_value
    
}


func generate_z_data( )string{
    
    
    
    return_value := `<div id="z_panel" style="width:100%;height:300pt;"></div>`
   
    return_value = return_value + `
    <script>
    `+
    `json_time_trace="`+ string(json_time_trace)+`";`+`
    
     json_z_trace="`+ string(json_z_trace)+`";`+`
     `+

    
    `   
    var time_trace         =  JSON.parse(json_time_trace)
    var z_trace            =  JSON.parse(json_z_trace)
    

    
    for( i = 0;i< time_trace.length; i++)
    { time_trace[i] = new Date(time_trace[i]/1e6)
    } 
    



     
     
    var trace1 = {
        x: time_trace,
        y: z_trace,
        name: 'Z_SCORE',
        mode: 'lines+markers',
       type: 'scatter'
     };
    

     
    var data = [trace1];

    
    
    
	z_panel = document.getElementById('z_panel');
	Plotly.newPlot( z_panel,data );
    </script>  `
    
    return return_value
    
}

func generate_v_data( )string{
    
    
    
    return_value := `<div id="v_panel" style="width:100%;height:300pt;"></div>`
   
    return_value = return_value + `
    <script>
    `+
    `json_time_trace="`+ string(json_time_trace)+`";`+`
    
     json_v_trace="`+ string(json_velocity_trace)+`";`+`
     `+

    
    `   
    var time_trace         =  JSON.parse(json_time_trace)
    var v_trace            =  JSON.parse(json_v_trace)
    

    
    for( i = 0;i< time_trace.length; i++)
    { time_trace[i] = new Date(time_trace[i]/1e6)
    } 
    



     
     
    var trace1 = {
        x: time_trace,
        y: v_trace,
        name: 'VELOCITY',
        mode: 'lines+markers',
       type: 'scatter'
     };
    

     
    var data = [trace1];

    
    
    
	v_panel = document.getElementById('v_panel');
	Plotly.newPlot( v_panel,data );
    </script>  `
    
    return return_value
    
}

 
func generate_filtered_table_display()string{
    
    display_list := make([][]string,len(filtered_return_value)) 
       
    for index,data            := range filtered_return_value {
        current_value         := fmt.Sprintf("%f",data.value.current_value)
        filtered_value        := fmt.Sprintf("%f",data.value.filtered_value)
        velocity              := fmt.Sprintf("%f",data.value.current_velocity)
        std                   := fmt.Sprintf("%f",data.value.z_value)
        z_value               := fmt.Sprintf("%f",data.value.std)
        z_state               := fmt.Sprintf("%t",data.value.z_state  )
        time_stamp            :=  format_time(data.time_stamp, true) 
        display_list[index] = []string{current_value,filtered_value, velocity, std, z_value,z_state, time_stamp} 
    
    }
    return web_support.Setup_data_table("raw_list",[]string{"CURRENT VALUE","FILTERED_VALUE","VELOCITY","STD","Z_VALUE","Z_STATE",  "TIME" },display_list)
    
}
   

func generate_func_data(key string){
    
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
    pg_data,status := stream_control.filtered_data_stream.Select_where(where_clause)

 
    
 
    filtered_return_value = make([]filtered_return_value_type,len(pg_data))
    if status != true {
        panic("bad select")
    }
    time_trace       := make([]int64,len(pg_data))
    filter_trace     := make([]float64,len(pg_data))
    std_trace        := make([]float64,len(pg_data))
    raw_trace        := make([]float64,len(pg_data))
    velocity_trace   := make([]float64,len(pg_data))
    z_trace          := make([]float64,len(pg_data))
    z_state          := make([]bool,len(pg_data))
    for index,data_element := range pg_data{
       var item   map[string]interface{}
       
       input := data_element.Data
       err :=  msgpack.Unmarshal([]byte(input), &item)
       if err != nil {
           panic("bad packed data")
       }
      
      
      var value filtered_value_type
      
      value.filtered_value    = item["filtered_value"].(float64)
      value.current_value     = item["current_value"].(float64)
      value.current_velocity  = item["current_velocity"].(float64)
      value.z_value           = item["data_z"].(float64)
      value.std               = item["std"].(float64)
      value.z_state           = item["z_state"].(bool)
      
 
      filtered_return_value[index].value = value
      filtered_return_value[index].time_stamp = data_element.Time_stamp
      time_trace[index]            = data_element.Time_stamp
      filter_trace[index]          = value.filtered_value
      std_trace[index]             = value.std
      raw_trace[index]             = value.current_value
      z_trace[index]               = value.z_value
      velocity_trace[index]        = value.current_velocity
      z_state[index]               = value.z_state
    }
    temp,_                     :=  json.Marshal(time_trace)
    json_time_trace            =  string(temp)
    temp,_                     =  json.Marshal(filter_trace)
    json_filter_trace          =  string(temp)
    temp,_                     =  json.Marshal(std_trace)
    json_std_trace             =  string(temp)
    temp,_                     =  json.Marshal(raw_trace)
    json_raw_trace             =  string(temp)
    temp,_                     =  json.Marshal(z_trace)
    json_z_trace               =  string(temp)
    temp,_                     =  json.Marshal(velocity_trace)  
    json_velocity_trace        =  string(temp)
    temp,_                     =  json.Marshal(z_state)
    json_zstate_trace          =  string(temp)
    
    
}


