package irrigation_modules







type irrigation_job_type struct {
   command   string
   schdule   string
   step      int64
   active    bool
   io        map[string]int
   time      int64

}

var offline_job irrigation_job_type 



var command_channel  chan string



func Initialize_channels(){ // called by  main

   command_channel = make(chan string,20)
   offline_job.active = false


}




func translate_offline_job( input map[string]interface{} )bool{

   if offline_job.active == true {
      return false
   }
   // do the unpacking
   return true

}


func send_Command_Channel( input string ){

    command_channel <- input


}

func get_Command_Channel()( string, bool){
  if len(command_channel) != 0 {
     return <-command_channel, true
  }
  return "",false


}


func decode_channel_command_field( input  *map[string]interface{})string{


   return ""
}




func decode_channel_io_fields(input  *map[string]interface{})[]map[string]int{

   var return_value = make([]map[string]int,0)
   
   return return_value
}

func decode_channel_time_field(input  *map[string]interface{})int{


   return 0
}