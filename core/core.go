package core

import (
//	"./models"
	"./webserver"
)

func Run() {

//	models.ResetDb()

	webserver.Run("127.0.0.1:8080")

	for {}
}