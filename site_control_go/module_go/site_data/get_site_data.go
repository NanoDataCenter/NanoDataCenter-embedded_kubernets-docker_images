package get_site_data

import ( 
"fmt"
"io/ioutil"
"time"
"encoding/json"


)

func Can_Not_Continue(display_string string){

   
   var delay_count = time.Second*10
   for{
     fmt.Println(display_string)
	 time.Sleep(delay_count)
   }
   
}

func Get_site_data(file_name string) map[string]interface{} {

         
    var data, err = ioutil.ReadFile(file_name)
    if err != nil {
        Can_Not_Continue("Bad File")
    }

    var site_data  map[string]interface{}
	var err1 = json.Unmarshal(data,&site_data)
    if err1 != nil{
	  Can_Not_Continue("bad json data")
	}
	return site_data
}



func Determine_master(site_file string) map[string]interface{} {

       var site_data = Get_site_data(site_file)
	   var val,ok = site_data["master"]
	   if ok&&val != true {
	       Can_Not_Continue("Not Master -- Spining in loop")
	   }
	   
       return site_data  
       
}
	  





