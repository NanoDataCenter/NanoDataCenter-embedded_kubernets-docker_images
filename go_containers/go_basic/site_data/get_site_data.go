package get_site_data

import ( 
//"fmt"
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



	  





