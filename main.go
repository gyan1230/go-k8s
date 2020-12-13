package main

import (
	"fmt"

	app "github.com/gyan1230/go-k8s/application"
	// app "github.com/gyan1230/2020/docker-kubernetes/go-k8s/application"
)

func main() {
	fmt.Println("Starting main go routine:::")
	app.Start()
}
