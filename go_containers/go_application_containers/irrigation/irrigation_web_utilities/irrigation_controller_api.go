package irrigation_web_support


func Queue_irrigation_jobs()string{
    return_value := `
    
    function queue_irrigation_direct(station,io){
        data = {}
        data["station"] = station
        data["io"]          = io
        ajax_post_confirmation("ajax/irrigation/irrigation_manage/post_job", data, "schedule station "+station +"  io  "+io, 
                                                  "action queued", "action not queued" )
}
    
    
    `
    return return_value
}
