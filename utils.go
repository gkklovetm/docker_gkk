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

func Go(f func() error) chan error {
	ch := make(chan error)
	go func() {
		ch <- f()
	}()
	return ch
}

func SelfPath() string {
	path, err := exec.LookPath(os.Args[0])
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
