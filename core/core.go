package core

import (
//	"./models"
	"./webserver"
//	"./parser"
//	"fmt"
)

func Run() {

//	models.ResetDb()
//
//	quitChan, totalChan, statusChan, errorChan := parser.Run(0, 0)
//
//	var total int
//	for  {
//		select {
//		case current := <- statusChan:
//			fmt.Printf("%d/%d\n", total, current)
//		case t := <- totalChan:
//			total = t
//		case err := <- errorChan:
//			fmt.Println(err)
//		case <- quitChan:
//			return
//
//		default:
//
//		}
//	}

	webserver.Run("127.0.0.1:8080")

	for {}
}