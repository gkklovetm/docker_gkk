package docker

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
)

func p() {
   fmt.Println("1")
}

func SelfPath() string{
   path, err :=  exec.LookPath(os.Args[0])
   if err != nil {
       panic(err)
   }
   path, err = filepath.Abs(path)
   if err != nil {
       panic(err)
   }
   fmt.Println(path)
   return path
}

