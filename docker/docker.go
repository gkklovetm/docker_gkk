package main


import (
    docker "../"
    "fmt"
)


func p() {
    fmt.Println("123")
   
}


func main() {
    if docker.SelfPath() == "/sbin/init" {
           docker.SysInit()
           return 
    } 
}
