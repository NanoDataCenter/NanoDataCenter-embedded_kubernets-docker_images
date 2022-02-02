package web_support

func Load_jquery_ajax_components()string {

return_value := `

function deepCopyObject(input)
{
  return JSON.parse(JSON.stringify(input))
}

function set_status_bar( text )
{
   $("#status_display").text(text)
}


var user_function 
function status_successful( data )
{
   set_status_bar("Current Status: Fetch operation successful")
   user_function(data)
}

function ajax_get( url_path, error_message, success_function )
{
   user_function = success_function
	  $("#status_display").text("Current Status: Operation in Progress")
   $.ajax(
   {
       type: "GET",
       url: url_path,
       dataType: 'json',
       async: true,
       //json object to sent to the authentication url
       success: status_successful,
              
       error: function () 
		    {
           set_status_bar("Current Status: "+error_message)  
		       
		       
      }
   });
}

function ajax_post_confirmation(url_path, data, confirmation_string, 
                                       success_message, error_message )
{
 
   var result = confirm(confirmation_string);  // change this
   if( result == true )
   {
       $("#status_display").text("Current Status: Operation in Progress")

       var json_string = JSON.stringify(data);
       $.ajax ({  type: "POST",
                  url: url_path,
                  dataType: 'json',
	                 contentType: "application/json",
                  async: true,
                  data: json_string,
                  success: function () 
		                {
                       set_status_bar(success_message)

		                 },

                   error: function () 
		                {
                       set_status_bar(error_message) 	                
                  }
           })
   }
}


function ajax_post(url_path, data,  success_message, error_message )
{
 
   var result = true
   if( result == true )
   {
       $("#status_display").text("Current Status: Operation in Progress")

       var json_string = JSON.stringify(data);
       $.ajax ({  type: "POST",
                  url: url_path,
                  dataType: 'json',
	                 contentType: "application/json",
                  async: true,
                  data: json_string,
                  success: function () 
		                {
                       set_status_bar(success_message)

		                 },

                   error: function () 
		                {
                       set_status_bar(error_message) 	                
                  }
           })
   }
}

function ajax_post_get(url_path, data, success_function, error_message) 
{
     var json_string = JSON.stringify(data);
     $("#status_display").text("Current Status: Operation in Progress")
     user_function = success_function

     $.ajax ({  type: "POST",
                  url: url_path,
                  dataType: 'json',
	                 contentType: "application/json",
                  async: true,
                  data: json_string,
                  success: status_successful,

                   error: function () 
		                {
                       set_status_bar("Current Status: "+error_message)  
		                 }
           })
   
}

`
 
return return_value
}

func Check_box_state_components()string{
    
    js := `
   
  function check_state(keys){
  
   let check_status = [];
   //console.log(keys)
   for (i= 0;i<keys.length;i++){
      let key = keys[i]
      //console.log("loop")
      //console.log(i)
      //console.log(key)
      if( $("#"+key).is(":checked") == true )
      {       
	         check_status.push(key);
	         
       }
        
    }
   return check_status
     
  }
 
  
  function select_values(keys){

  let i = 0
  for( i= 0;i<keys.length;i++){
      let key = keys[i]
      $("#"+key).prop('checked', true)
   }

}

function unselect_values(keys){
  let i = 0
  for( i= 0;i<keys.length;i++){
      let key = keys[i]
      $("#"+key).prop('checked', false)
   }


}`

return js
}




