package main

import (
	docker ".."
	"../rcli"
	"flag"
	"fmt"
	"log"
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
	fmt.Println("tm")
	return nil
}
