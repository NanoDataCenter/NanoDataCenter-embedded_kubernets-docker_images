package irrigation_web_support


func Queue_irrigation_jobs()string{
    return_value := `
    var  direct_io_state 
    var direct_post_station 
    var direct_post_io
    var direct_post_time
    var direct_post_message
    
     var current_post_station 
    var   current_post_io
   
function initialize_direct_io_control(){
     direct_io_state = false
}
 
    
function queue_irrigation_direct(station,io,time,message){
        if( direct_io_state == true)
        {
           clear_then_queue_irrigation_direct(station,io,time,message)
           return
         }
         direct_io_state = true
        let data = {}
        data["station"] = station
        data["io"]          = io
        data["time"]     = time
        data["action"] = true
        current_post_station = station 
        current_post_io           = io

        ajax_post_confirmation("ajax/irrigation/irrigation_manage/post_job", data, message, 
                                                  "action queued", "action not queued" )
}


function   clear_then_queue_irrigation_direct(station,io,time,message){
       if( confirm(message) == false){
           return
       }
        direct_post_io = io
        direct_post_time = time
        direct_post_station = station
         
        let data = {}
        data["station"] = current_post_station 
        data["io"]          = current_post_io
        data["time"]     = 1 // default place holder
        data["action"] = false
      
        ajax_post_get_new("ajax/irrigation/irrigation_manage/post_job", data, post_current_job, "job not cleared") 
    
    
   }

function post_current_job(){
       data = {}
        data["station"] = direct_post_station
        data["io"]          = direct_post_io
        data["time"]     = direct_post_time
        data["action"] = true
         current_post_station = direct_post_station
        current_post_io           = direct_post_io
         
        ajax_post("ajax/irrigation/irrigation_manage/post_job", data,  
                                                  "action queued", "action not queued" )
}      



function direct_io_close_then_action( action){
       if( direct_io_state == false ){
           action()
           return
        }
       
        
         
        data = {}
        data["station"] = current_post_station 
        data["io"]          = current_post_io
        data["time"]     = 1 // default place holder
        data["action"] = false
       ajax_post_get("ajax/irrigation/irrigation_manage/post_job", data, action, "job not cleared") 
        
    
   }
    `
    return return_value
}
