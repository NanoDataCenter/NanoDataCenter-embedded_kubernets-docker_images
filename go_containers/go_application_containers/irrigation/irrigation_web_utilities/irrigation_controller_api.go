package irrigation_web_support


func Queue_irrigation_jobs()string{
    return_value := `
    
    function queue_irrigation_direct(station,io,time,message){
        data = {}
        data["station"] = station
        data["io"]          = io
        data["time"]     = time
        ajax_post_confirmation("ajax/irrigation/irrigation_manage/post_job", data, message, 
                                                  "action queued", "action not queued" )
}
    
    
    `
    return return_value
}
