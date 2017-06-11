package docker

import (
	auth "./auth"
	"container/list"
)

type Runtime struct {
	root       string
	repository string
	containers *list.List
	authConfig *auth.AuthConfig
}

func NewRuntime() (*Runtime, error) {
	return NewRuntime
}

func NewRuntime() (*Runtime, error) {
	//return NewRuntimeFromDirectory(root:"/var/lib/docker")
	runtime := &Runtime{
		root:       "gkk",
		repository: "tm",
		containers: list.New(),
		authConfig: auth.NewAuthConfig("krunerge", "gkk82951173", "1056483357@qq.com", "/root"),
	}
	return runtime, nil
}
