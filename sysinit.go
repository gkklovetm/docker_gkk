package docker

import (
    "os"
    "flag"
    "fmt"
)


func SysInit() {
    if len(os.Args) <= 1 {
        fmt.Println("You should not invoke docker-init manually")
        os.Exit(1)
    }
    var u = flag.String("u", "", "username or uid")
    var gw = flag.String("g", "", "gateway address")
   
    flag.Parse()

   // setupNetworking(*gw)
    fmt.Println("%s", *u)
    fmt.Println("%s", *gw) 
}
