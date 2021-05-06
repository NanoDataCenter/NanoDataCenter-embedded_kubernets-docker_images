package main


import (
    //"fmt"
	"strconv"
    "net/http"
	"lacima.com/Patterns/web_server"
	"lacima.com/site_data"
    "lacima.com/redis_support/graph_query"
    "lacima.com/redis_support/redis_handlers"
    "lacima.com/redis_support/generate_handlers"
)




func main() {

 var config_file = "/data/redis_server.json"
  var site_data_store map[string]interface{}

  site_data_store = get_site_data.Get_site_data(config_file)
  graph_query.Graph_support_init(&site_data_store)
  redis_handlers.Init_Redis_Mutex()
  data_handler.Data_handler_init(&site_data_store)
  
  
  port, ok_flag := web_support.Setup_Web_Server( "irrigation")
  if ok_flag != true {
     panic("web server not registered")
  }
  initialize_handlers()
 
  if err := http.ListenAndServe(":"+strconv.Itoa(int(port)), nil); err != nil {
        panic(err)
  }
}

func initialize_handlers(){

 fileServer := http.FileServer(http.Dir("./static"))
 http.Handle("/", fileServer)






}




/*
  if err = http.ListenAndServe(":80", nil); err != nil {
        panic(err)
  }
   //fileServer := http.FileServer(http.Dir("./static"))
    //http.Handle("/", fileServer)
 
   http.HandleFunc("/form", formHandler)
    http.HandleFunc("/hello", helloHandler)


func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "POST request successful")
    name := r.FormValue("name")
    address := r.FormValue("address")
    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    //if r.URL.Path != "/hello" {
    //    http.Error(w, "404 not found.", http.StatusNotFound)
    //    return
    //}

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }


    fmt.Fprintf(w, "Hello!")
}
*/