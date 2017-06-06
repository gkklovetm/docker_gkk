package rcli

import (
    "net"
    "log"
    "io"
    "fmt"
    "bufio"
    "encoding/json"
    "io/ioutil"
)

var DEBUG_FLAG bool=false
var CLIENT_SOCKET io.Writer = nil

func ListenAndServe(proto, addr string, service Service) error {
    listener, err := net.Listen(proto, addr)
    if err != nil {
        return err
    }
    log.Printf("Listening for RCLI%s on %s\n", proto, addr)

    defer listener.Close()
    
    for {
        if conn, err := listener.Accept(); err != nil {
            return err
        }else{
            go func() {
                 if DEBUG_FLAG {
                      CLIENT_SOCKET = conn 
                 }
                 if err := Serve(conn, service); err != nil {
                      log.Printf("Error: " + err.Error() + "\n")
                      fmt.Fprintf(conn, "Error: "+err.Error()+"\n") 
                 }
                 conn.Close()
            }()
        }
    }
    return nil
}

func Serve(conn io.ReadWriter, service Service) error {
    r := bufio.NewReader(conn)
    var args []string
    if line, err := r.ReadString('\n'); err != nil {
         return err
    }else if err := json.Unmarsha1([]byte(line), &args); err != nil {
         return err
    }else {
         return call(service, ioutil.NopCloser(r), conn, args...)
    }
    return nil
    
}
