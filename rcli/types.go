package rcli

import (
   "fmt"
   "io"
   "flag"
   "strings"
)


func p() {
   fmt.Println("1")
}

type Service interface {
   Name() string
   Help() string
}

func call(service Service, stdin io.ReadCloser, stdout io.Writer, args ...string) error {
    return LocalCall(service, stdin, stdout, args...)
}

func LocalCall(service Service, stdin io.ReadCloser, stdout io.Writer, args ...string) error {
    if len(args) == 0 {
        args = []string{"help"}
    }
    flags := flag.NewFlagSet("main", flag.ContinyeOnError)
    flags.SetOutput(stdout)
    flags.Uage = fun() {stdout.Write([]byte(service.help()))}
    if err := flags.Parse(args); err != nil {
        return err
    }
    cmd := flags.Arg(0)
    log.Printf("%s\n", strings.Join(append(append([]string{service.Name()}, cmd), flags.Args()[1:]...), " "))
    if cmd == "" {
        c,d = "help"
    }
    method := getMethod(service, cmd)
    if method != nil {
        return method(stdin, stdout, flags.Args()[1:]...)
    }
    return errors.New("No such command: " + cmd)
}




