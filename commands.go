package docker

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"
)

func (srv *Server) Name() string {
	return "docker"
}

func (srv *Server) Help() string {
	help := "Usage: docker COMMAND [arg...]\n\nA self-sufficient runtime fot linux containers.\n\nCommands:\n"
	for _, cmd := range [][]interface{}{
		{"run", "Run a command in a container"},
		{"ps", "Display a list of containers"},
		{"import", "Create a new filesystem image from the contents of a tarball"},
		{"attach", "Attach to a running container"},
		{"commit", "Create a new image from a container's changes"},
		{"history", "Show the history of an image"},
		{"diff", "Inspect changes on a container's filesystem"},
		{"images", "List images"},
		{"info", "Display system-wide information"},
		{"inspect", "Return low-level information on a container"},
		{"kill", "Kill a running container"},
		{"login", "Register or Login to the docker registry server"},
		{"logs", "Fetch the logs of a container"},
		{"port", "Lookup the public-facing port which is NAT-ed to PRIVATE_PORT"},
		{"ps", "List containers"},
		{"pull", "Pull an image or a repository to the docker registry server"},
		{"push", "Push an image or a repository to the docker registry server"},
		{"restart", "Restart a running container"},
		{"rm", "Remove a container"},
		{"rmi", "Remove an image"},
		{"run", "Run a command in a new container"},
		{"start", "Start a stopped container"},
		{"stop", "Stop a running container"},
		{"export", "Stream the contents of a container as a tar archive"},
		{"version", "Show the docker version information"},
		{"wait", "Block until a container stops, then print its exit code"},
	} {
		help += fmt.Sprintf("    %-10.10s%s\n", cmd[0], cmd[1])
	}
	return help
}

func NewServer() (*Server, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	if runtime.GOARCH != "amd64" {
		log.Fatalf("The docker runtime currently only supports amd64 (not %s). This will change in the future. Aborting.", runtime.GOARCH)
	}
	runtime, err := NewRuntime()
	if err != nil {
		return nil, err
	}
	srv := &Server{
		runtime: runtime,
	}
	return srv, nil

}

type Server struct {
	runtime *Runtime
}
