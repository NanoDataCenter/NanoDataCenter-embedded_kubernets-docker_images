package mqtt_topics

import	"github.com/msgpack/msgpack-go"




// returns msgpack data


type Topic_generator_type func( data interface {} )string

var topic_map[string]Topic_generator_type


func Init_topic_generation(){
  topic_map = make(topic_map[string]interface{})
  
  topic_map["string"]          = generate_string_msgpack
  topic_map["int8_array"]      = generate_int8_msgpack
  topic_map["int16_array"]     = generate_int16_msgpack
  topic_map["int32_array"]     = generate_int32_msgpack
  topic_map["int64_array"]     = generate_int64_msgpack
  topic_map["uint8_array"]     = generate_uint8_msgpack
  topic_map["uint16_array"]    = generate_uint16_msgpack
  topic_map["uint32_array"]    = generate_uint32_msgpack 
  topic_map["uint64_array"]    = generate_uint64_msgpack  
  topic_map["float32_array"]   = generate_float32_msgpack
  topic_map["float64_array"]   = generate_float64_msgpack
  topic_map["map"]             = generate_map
  
}

func Register_topic_generator( id string typic_generator Topic_generator_type){
    
    if _,ok := topic_map[topic_type] ; ok == false {
       panic("bad topic type")
   }
   topic_data = topic_map[topic_type](data)   
    
    
    
}

func Generate_topic_data( topic_type string, data interface{} )string {  

   if _,ok := topic_map[topic_type] ; ok == false {
       panic("bad topic type")
   }
   topic_data = topic_map[topic_type](data)
}



func generate_string_msgpack( input interface{})string{
   data := input.(string)
   var b1 bytes.Buffer	
   msgpack.Pack(&b1,time_stamp)
   return b1.String()
    
    
}

func  generate_int8_msgpack( input interface{})string{
   data := input.([]int8)
   var b1 bytes.Buffer	
   msgpack.Pack(&b1,time_stamp)
   return b1.String()
    
    
}    
    
    
func  generate_int16_msgpack( input interface{})string{
   data := input.([]int8)
   var b1 bytes.Buffer	
   msgpack.Pack(&b1,time_stamp)
   return b1.String()
    
    
}

func  generate_int32_msgpack( input interface{})string{
   data := input.([]int8)
   var b1 bytes.Buffer	
   msgpack.Pack(&b1,time_stamp)
   return b1.String()
    
    
}

func  generate_int64_msgpack( input interface{})string{
   data := input.([]int64)
   var b1 bytes.Buffer	
   msgpack.Pack(&b1,time_stamp)
   return b1.String()
    
    
}

func  generate_uint8_msgpack( input interface{})string{
   data := input.([]uint8)
   var b1 bytes.Buffer	
   msgpack.Pack(&b1,time_stamp)
   return b1.String()
    
    
}

func  generate_uint16_msgpack( input interface{})string{
   data := input.([]uint16)
   var b1 bytes.Buffer	
   msgpack.Pack(&b1,time_stamp)
   return b1.String()
    
    
}

func  generate_uint32_msgpack( input interface{})string{
   data := input.([]uint32)
   var b1 bytes.Buffer	
   msgpack.Pack(&b1,time_stamp)
   return b1.String()
    
    
}

func  generate_uint64_msgpack( input interface{})string{
   data := input.([]uint64)
   var b1 bytes.Buffer	
   msgpack.Pack(&b1,time_stamp)
   return b1.String()
    
    
}

func  generate_float32_msgpack( input interface{})string{
   data := input.([]float32)
   var b1 bytes.Buffer	
   msgpack.Pack(&b1,time_stamp)
   return b1.String()
    
    
}

func  generate_float64_msgpack( input interface{})string{
   data := input.([]float64)
   var b1 bytes.Buffer	
   msgpack.Pack(&b1,time_stamp)
   return b1.String()
    
    
}

func  generate_map( input interface{})string{
   data := input.(map[string]interface{})
   var b1 bytes.Buffer	
   msgpack.Pack(&b1,time_stamp)
   return b1.String()
    
    
}
