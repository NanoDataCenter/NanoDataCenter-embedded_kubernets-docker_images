package main

import (
	"fmt"
    "example.com/user/hello/docker_control"
)

func test_prune(){
  docker_control.Prune()
}


func test_images() {
    var images = docker_control.Images()
	
	for i:=0;i<len(images);i++{
	  fmt.Println(images[i] )
	}

}


func test_container_all(){

   var container_list = docker_control.Containers_ls_all()
   fmt.Println(container_list)
   var container_list_1 = docker_control.Containers_ls_runing()
   fmt.Println(container_list_1)

}

func main(){
  //test_images()
  //test_prune()
  //test_images()
  test_container_all()
}