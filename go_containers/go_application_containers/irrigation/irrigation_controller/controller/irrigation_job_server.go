package irrigation_controller

import (
      "fmt"
      "time"
      
    
)

func Start() {
	
    fmt.Println("job server  initialization")
	for true {
        fmt.Println("job server polling")
		time.Sleep(time.Second * 60)
	

	}

}
