package main

import (
	docker ".."
	"../rcli"
	"../term"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func p() {
	fmt.Println("123")

}

func main() {
	if docker.SelfPath() == "/sbin/init" {
		docker.SysInit()
		return
	}

	fl_daemon := flag.Bool("d", false, "Daemon mode")
	fl_debug := flag.Bool("D", false, "Debug mode")

	flag.Parse()

	rcli.DEBUG_FLAG = *fl_debug
	if *fl_daemon {
		if flag.NArg() != 0 {
			flag.Usage()
			return
		}
		if err := daemon(); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := runCommand(flag.Args()); err != nil {
			log.Fatal(err)
		}
	}
}

func daemon() error {
	//p()
	fmt.Println("ggk")
	//return nil
	service, err := docker.NewServer()
	if err != nil {
		return err
	}

	return rcli.ListenAndServe("tcp", "127.0.0.1:4545", service)
}

func runCommand(args []string) error {
	p()
	var oldState *term.State
	var err error
	if term.IsTerminal(0) && os.Getenv("NORAW") == "" {
		oldState, err = term.MakeRaw(0)
		if err != nil {
			return err
		}
		defer term.Restore(0, oldState)
	}
	if conn, err := rcli.Call("tcp", "127.0.0.1:4545", args...); err == nil {
		receive_stdout := docker.Go(func() error {
			_, err := io.Copy(os.Stdout, conn)
			return err
		})
		send_stdin := docker.Go(func() error {
			_, err := io.Copy(conn, os.Stdin)
			if err := conn.CloseWrite(); err != nil {
				log.Printf("Couldn't send EOF: " + err.Error())
			}
			return err
		})
		if err := <-receive_stdout; err != nil {
			return err
		}
		if !term.IsTerminal(0) {
			if err := <-send_stdin; err != nil {
				return err
			}
		}

	} else {
		service, err := docker.NewServer()
		log.Printf("docker server run local")
		if err != nil {
			return err
		}
		if err := rcli.LocalCall(service, os.Stdin, os.Stdout, args...); err != nil {
			return err
		}

	}
	if oldState != nil {
		term.Restore(0, oldState)
	}
	return nil
}
