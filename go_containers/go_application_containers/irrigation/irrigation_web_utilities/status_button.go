package irrigation_web_support


func Attach_status_panel()string{
 
    return_value := `
    
    <script type="text/javascript" >
  $(document).ready(function(){
    $("#status_panel").click(function(){
    
        $("#status_modal").modal("show");
        status_request_function()
    });
    
});
     
 



function status_request_function()
{
   json_object = {}

   
   //ajax_get( '/ajax/status_update', "Data Not Fetched", status_update )
   
   
}
  
                  
function status_update( data )
{
      return
/*
       var temp
       var temp_1
       var tempDate
       
       
       //alert(JSON.stringify(data))
       //alert(Object.keys(data))
       var date = new Date( data["TIME_STAMP"]  * 1000);
       tempDate = new Date()
       
       $("#time_stamp").html("Time:  "+tempDate.toLocaleDateString() + "   " + tempDate.toLocaleTimeString() )

       $("#controller_time_stamp").html("Ctr Time: "+ date.toLocaleDateString() + "   " + date.toLocaleTimeString() )
       $("#flow_rate").html("Current Flow Rate: "+parseFloat(data.MAIN_FLOW_METER).toFixed(2));
       $("#plc_flow_rate").html("PLC Flow Rate: "+parseFloat(data.PLC_FLOW_METER).toFixed(2));
       $("#cleaning_rate").html("Cleaning Flow Rate:  "+parseFloat(data.CLEANING_FLOW_METER).toFixed(2));
       $("#schedule").html("Schedule: "+data["SCHEDULE_NAME"])
       $("#step").html("Step:   "+data["STEP"])
       $("#time").html("Step Time:  "+data["RUN_TIME"])
       $("#duration").html("ELASPED_TIME: "+ data["ELASPED_TIME"]) 
       $("#rain_day").html("Rain Day: "+data["RAIN_FLAG"])
       $("#irrigation_current").html("Irrigation  Current:"+parseFloat(data.PLC_IRRIGATION_CURRENT).toFixed(2))
       $("#equipment_current").html("Equipment  Current:  "+parseFloat(data.PLC_EQUIPMENT_CURRENT).toFixed(2))
       $("#pump_input_current").html("Pump Input Current: "+parseFloat(data.INPUT_PUMP_CURRENT).toFixed(2))
       $("#pump_output_current").html("Pump Output Current: "+parseFloat(data.OUTPUT_PUMP_CURRENT).toFixed(2))
       $("#well_pressure").html("Well Pressure:  "+parseFloat(data.WELL_PRESSURE).toFixed(2))
       $("#master_valve").html("Master Valve: "+data.MASTER_VALVE )
       $("#eto_management").html("ETO Management: "+data.ETO_MANAGEMENT )
 
       $("#suspend").html("Suspension State:  "+data.SUSPEND )
       $("#clean_filter_limit").html("Filter Cleaning Limit (Gallon):  "+parseFloat(data.CLEANING_INTERVAL).toFixed(2) )
       $("#clean_filter_value").html("Filter Cleaning Accumulation (GPM):  "+parseFloat(data.CLEANING_ACCUMULATION).toFixed(2 ))
*/

}





    </script>
    
    `
    
    return_value = return_value + `
   
<div class="modal fade" id="status_modal" tabindex="-1" role="dialog" aria-labelledby="accountModalLabel" aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="accountModalLabel">Irrigation State</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="close">
                    <span aria-hidden="true">&times;</span>
                </button>
                 
            </div>
                 <ul >
      
                   <li id="time_stamp">Time Stamp: </li>
                   <li id="controller_time_stamp">Controller Time Stamp: </li>
                   <li id="schedule">Sprinkler Schedule: </li>
                   <li id="step">Sprinkler Step:   </li>
                   <li id="time">Time Of Step:      </li>
                   <li id="duration">Current Duration:  </li>  
                   <li id="flow_rate">Current Flow Rate:  </li>
                   <li id="plc_flow_rate">PLC Flow Rate:  </li>
                   <li id="cleaning_rate">Cleaning Flow Rate:  </li>
                   <li id="irrigation_current">Irrigation  Current: </li>
                   <li id="equipment_current">Equipment  Current: </li>
                   <li id="pump_input_current">Pump Input Current: </li>
                   <li id="pump_output_current">Pump Output Current: </li>
                   <li id="well_pressure">Well Pressure;   </li>
                   <li id="rain_day">Rain Day:      </li>        
                   <li id="eto_management">ETO Management: </li>
                   <li id="master_valve">Master Valve State: </li>
                    <li id="suspend">SUSPENSION STATE: </li>
                   <li id="clean_filter_limit">Filter Cleaning Limit (GPM): </li>
                  <li id="clean_filter_value">Filter Cleaning Accumulation (GPM): </li>
                </ul> 
                <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
                    
                 </div>
           </div>
    </div>
</div>   
    
    
    `
    return return_value
    
}


