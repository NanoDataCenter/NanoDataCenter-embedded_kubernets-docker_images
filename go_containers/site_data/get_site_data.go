package get_site_data

import ( 
"fmt"
"io/ioutil"
//"time"
"encoding/json"


)



func Get_site_data(file_name string) map[string]interface{} {

    
    var data, err = ioutil.ReadFile(file_name)
    if err != nil {
        panic("no configuration file")
    }

    var site_data  map[string]interface{}
	var err1 = json.Unmarshal(data,&site_data)
    if err1 != nil{
	 panic("bad json site data")
	}
	return site_data
}

func Save_site_data(file_name string, data map[string]interface{})  {
   //fmt.Println("***************************")
    //fmt.Println("file_name",file_name)
    //fmt.Println("data",data)
    json_data, _ := json.MarshalIndent(data,"","")
 
	err := ioutil.WriteFile(file_name, json_data, 0644)
    if err != nil{
        fmt.Println(err)
        panic("bad  json write")
    }

}




	  





