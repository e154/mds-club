package core

import (
	"./models"
	"./parser"
)

func Run() {

	models.ResetDb()
	parser.Run()
}