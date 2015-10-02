package core

import (
//	"./models"
	"./webserver"
	_ "./parser"
//	"fmt"
)

func Run() {

	webserver.Run("127.0.0.1:8080")

	for {}
}