package docker

import (
	"./auth"
	"container/list"
)

type Runtime struct {
	root       string
	repository string
	containers *list.List
	authConfig *auth.AuthConfig
}

func NewRuntime() (*Runtime, error) {
	//return NewRuntimeFromDirectory(root:"/var/lib/docker")

	runtime := &Runtime{
		root:       "gkk",
		repository: "tm",
		containers: list.New(),
	}
	return runtime, nil
}
