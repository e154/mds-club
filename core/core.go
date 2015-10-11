package core

import (
//	"./models"
	"./webserver"
	_ "./parser"
	"fmt"
)

const (
	major int = 0
	minor int = 1
	patch int = 0
	address string = "127.0.0.1:8080"
)

func Run() {

	fmt.Printf("MDS client v%d.%d.%d\n", major, minor, patch)
	fmt.Printf("Start server on http://%s\n", address)

	webserver.Run(address)

	for {}
}