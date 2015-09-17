package parser

import (
	"testing"
//	"fmt"
)



func TestGetTotalElements(t *testing.T) {

	total, err := GetTotalElements(URL)
	if err != nil {
		t.Errorf("error: %s\n", err.Error())
	}

	if total == 0 {
		t.Errorf("error: return zero value\n")
	}
}

func TestParser(t *testing.T) {

//	quitChan, totalChan, statusChan, errorChan := Run(0, 2)
//
//	var total int
//	for  {
//		select {
//		case current := <- statusChan:
//			fmt.Printf("%d/%d\n", total, current)
//		case total := <- totalChan:
//			t.Logf("total: %d\n", total)
//		case err := <- errorChan:
//			t.Errorf("error: %s\n", err.Error())
//		case <- quitChan:
//			return
//
//		default:
//
//		}
//	}
}
