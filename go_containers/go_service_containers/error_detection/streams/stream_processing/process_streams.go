package monitor_streams

import "math"
import "time"
import "fmt"
import "lacima.com/Patterns/msgpack_2"
import "lacima.com/server_libraries/postgres"
import "github.com/vmihailenco/msgpack/v5"
import "github.com/montanaflynn/stats"



var Z_LEVEL  float64

type Median_Filter_Type struct {
 
    buffer_position int64
    buffer_limit    int64
    median_buffer   []float64
    filtered_value  float64
    current_value   float64
}


type Velocity_Type struct {
    
   previous_value    float64
   current_velocity  float64
   lag_velocity      float64
   r_value           float64
    
}

type Z_Type struct {
    z_value           float64
    std               float64
    z_state           bool  
    
    
}

type Stream_Processing_Type struct {
    median    Median_Filter_Type
    velocity  Velocity_Type
    z_data    Z_Type
}
    
    

func Process_functions(){
    
    initalize_stream_processing()
    for true {
       
       fmt.Println("sample_time",monitor_control.sample_time)
       time.Sleep(time.Duration(monitor_control.sample_time)* time.Second)
       process_stream_logs()
      
    
    }
    
}

func initalize_stream_processing(){
    
    Z_LEVEL = 3.0
    monitor_control.current_time = time.Now().UnixNano()   
    
}

func process_stream_logs(){
    
     for true {
         stream_data,err :=monitor_control.process_data_stream.Select_after_time_stamp_asc( monitor_control.sample_time)
         ts := time.Now().Unix()
         fmt.Println("ts    ",ts)
         if err != true {
           
           fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$ select err EE$$$$$$$$$$$$$$$$$$$$$$",ts)
           time.Sleep(time.Second*10)
           continue
         }
        
    
        fmt.Println("data stream length",len(stream_data))
        
       for _,data_element := range stream_data {
       key_string :=    pg_drv.Assemble_key(data_element)
       fmt.Println("key_string",key_string)
       value,err      :=    msg_pack_utils.Unpack_float64(data_element.Data)
       if err != true {
           panic("bad packed data")
       }
       analyize_data_element(key_string,value,data_element.Time_stamp,data_element)
         
        }
        break
     }
     monitor_control.current_time = time.Now().UnixNano()
}


func analyize_data_element(key_string string,value float64,time_stamp int64, data_element pg_drv.Stream_Output_Data_Record){
   previous_time_stamp := get_time_stamp(key_string)
   
   if previous_time_stamp < time_stamp {
      
       process_data(key_string ,value ,time_stamp,data_element )

   }
}


func process_data(key_string string,value float64,time_stamp int64,data_element pg_drv.Stream_Output_Data_Record){
    
    stream_processing_data := get_stream_processing_data(key_string,value)
    //fmt.Println("stream_processing_old",key_string,time_stamp,stream_processing_data)
    stream_processing_data = process_entry(stream_processing_data,value)
    fmt.Println("stream_processing_data_new",key_string,time_stamp,value,stream_processing_data)
    
    store_stream_processing_data(key_string,stream_processing_data ,time_stamp,data_element )
    

}

func process_entry(data Stream_Processing_Type,value float64)Stream_Processing_Type{
    
    data.median      = process_median(data.median,value)
    data.velocity    = process_velocity(data)
    data.z_data      = process_z_data(data)

    return data
    
    
    
}

/*
 * type Median_Filter_Type struct {
 
    buffer_position int64
    buffer_limit    int64
    median_buffer   [5]float64
    filtered_value  float64
    current_value   float64
}



*/

func process_median(data Median_Filter_Type,value float64)Median_Filter_Type{
    
   return_value                      := data
   return_value.current_value        =  .1*value +.9*return_value.current_value
   index                             := return_value.buffer_position
   limit                             := return_value.buffer_limit
   return_value.median_buffer[index] = return_value.current_value
   index                             = index + 1
   if index >= limit {
       index                         = 0
   }
   return_value.buffer_position      = index
 
   
   //median, err                       := stats.Median(return_value.median_buffer)
  // if err != nil {
      // panic("bad median")
   //}
   
   return_value.filtered_value        = return_value.current_value

return  return_value
    
    
    
}

/*





type Velocity_Type struct {
    
   previous_value    float64
   current_velocity  float64
   lag_velocity      float64
   r_value           float64
    
}
*/

func process_velocity(data Stream_Processing_Type)Velocity_Type{
 
    return_value                    := data.velocity
    return_value.current_velocity   = (1. - return_value.r_value)*(data.median.filtered_value-return_value.previous_value)+(return_value.r_value*return_value.lag_velocity)
    return_value.lag_velocity       = return_value.current_velocity
    return_value.previous_value     = data.median.filtered_value
 

    return return_value
    
    
    
}

/*
type Z_Type struct {
    z_value           float64
    std               float64
    z_state           bool  
    
    
}
*/


func process_z_data(data Stream_Processing_Type)Z_Type{
    
  var return_value Z_Type
    
   std, err                       := stats.StandardDeviation(data.median.median_buffer)
   if err != nil {
       panic("bad median")
   }
   
   return_value.std = std
   if return_value.std == 0 {
       return_value.z_state = false
       return_value.z_value = 0
       return return_value
   }
   
   return_value.z_value    = math.Abs((data.median.current_value-data.median.filtered_value)/return_value.std)
   
   if return_value.z_value > Z_LEVEL {
       return_value.z_state = true
   }else{
    
       return_value.z_state = false
       
   }
   

   return return_value
    
    
    
}




func get_time_stamp(key_string string)int64{
    
    packed_time := monitor_control.time_table.HGet(key_string)
    time, err := msg_pack_utils.Unpack_int64(packed_time)
    if err != true {
        time = 0
    }
    return time
    
}

func get_stream_processing_data(key_string string,input_value float64)Stream_Processing_Type{
    
    var return_value Stream_Processing_Type

    intermediate_data := make(map[string]interface{})
    packed_data := monitor_control.stream_table.HGet(key_string)
    //fmt.Println("packed_data",len(packed_data),packed_data)
    packed_byte := []byte(packed_data)
    err := msgpack.Unmarshal(packed_byte, &intermediate_data )
    //fmt.Println("err",err)
    if err != nil {
      
        //fmt.Println("initializing stream",input_value)
	    return_value = initialize_stream_processing_data(input_value)
        
    }else{
      return_value = recover_intermediate_values(intermediate_data)
      //fmt.Println("recovered value",return_value)
     
    }
    
    //fmt.Println(return_value)    
    
    return return_value
    
    
}


 

func recover_intermediate_values( data map[string]interface{})Stream_Processing_Type{
 
    var return_value Stream_Processing_Type
    
    return_value.z_data.z_value                  = data["data_z"].(float64) 
    return_value.z_data.std                      = data["std"].(float64)
    return_value.z_data.z_state                  = data["z_state"].(bool)
    return_value.velocity.previous_value         = data["previous_value"].(float64)
    return_value.velocity.current_velocity       = data["current_velocity"].(float64)   
    return_value.velocity.lag_velocity           = data["lag_velocity"].(float64)  
    return_value.velocity.r_value                = data["r_value"].(float64)
    return_value.median.current_value            = data["current_value"].(float64)  
    return_value.median.buffer_position          = data["buffer_position"].(int64)
    return_value.median.buffer_limit             = data["buffer_limit"].(int64)  
    return_value.median.filtered_value           = data["filtered_value"].(float64)
    temp_buffer                                  := data["median_buffer"].([]interface{})
    return_value.median.median_buffer            = make([]float64,len(temp_buffer))
    for index, value := range temp_buffer{
        return_value.median.median_buffer[index] = value.(float64)
    }
    return return_value
    
    
}



func initialize_stream_processing_data( input_data float64)Stream_Processing_Type{
    var return_value Stream_Processing_Type
    return_value.z_data.z_value = 0.
    return_value.z_data.std     = 0.
    return_value.z_data.z_state = false
    return_value.velocity.previous_value            = input_data
    return_value.velocity.current_velocity = 0.
    return_value.velocity.lag_velocity     = 0.
    return_value.velocity.r_value          = .5
    return_value.median.current_value      = input_data
    return_value.median.buffer_position    = 1
    return_value.median.buffer_limit       = 5
    return_value.median.filtered_value     = input_data
    return_value.median.median_buffer      = make([]float64,5)
    return_value.median.median_buffer[0]   = input_data
    return_value.median.median_buffer[1]   = input_data
    return_value.median.median_buffer[2]   = input_data
    return_value.median.median_buffer[3]   = input_data
    return_value.median.median_buffer[4]   = input_data

    return return_value
}



func translate_data( data Stream_Processing_Type)map[string]interface{}{
  
    return_value := make(map[string]interface{})
    
    return_value["data_z"]             = data.z_data.z_value
    return_value["std"]                = data.z_data.std     
    return_value["z_state"]            = data.z_data.z_state 
    return_value["previous_value"]     = data.velocity.previous_value 
    return_value["current_velocity"]   = data.velocity.current_velocity 
    return_value["lag_velocity"]       = data.velocity.lag_velocity     
    return_value["r_value"]            = data.velocity.r_value         
    return_value["current_value"]      = data.median.current_value   
    return_value["buffer_position"]    = data.median.buffer_position  
    return_value["buffer_limit"]       = data.median.buffer_limit   
    return_value["filtered_value"]     = data.median.filtered_value  
    return_value["median_buffer"]      = data.median.median_buffer
    return return_value
    
}

func store_stream_processing_data(key_string string, data Stream_Processing_Type,time_stamp int64,data_element pg_drv.Stream_Output_Data_Record ){
     //fmt.Println("processed data",key_string,data)
     pre_packdata := translate_data( data )
     
     
     packed_time := msg_pack_utils.Pack_int64(time_stamp)
     packed_data, err := msgpack.Marshal(&pre_packdata)
     //fmt.Println("packed_data",len(packed_data),packed_data)
     if err != nil {
	    fmt.Println("bad stored data")
        panic("bad data")
     }
    
    
    
     monitor_control.stream_table.HSet(key_string,string(packed_data))
     monitor_control.time_table.HSet(key_string,string(packed_time))
     fmt.Println(monitor_control.filtered_data_stream.Insert(data_element.Tag1,data_element.Tag2,data_element.Tag3,data_element.Tag4,data_element.Tag5,string(packed_data)))
     
     
     if data.z_data.z_state == true {
         fmt.Println("zstate true",data)
         monitor_control.z_table.HSet(key_string,string(packed_data))
         monitor_control.z_time.HSet(key_string,string(packed_time))
         monitor_control.process_incident_stream.Insert(data_element.Tag1,data_element.Tag2,data_element.Tag3,data_element.Tag4,data_element.Tag5,string(packed_data))
     }
    
}




func store_working_data(key string,working_data Stream_Processing_Type,time_stamp int64){
    monitor_control.time_table.HSet(key,msg_pack_utils.Pack_int64(time_stamp))

    
}
