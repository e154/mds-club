package core

import (
	"testing"
	"fmt"
)

func TestDownload(t *testing.T) {

	quitChan, dataChan, errChan := Download("http://mds.kallisto.ru/pionerfm/Dzherom_K._Dzherom_-_Kot_Dika_Dankermana.mp3")

	for {
		select {
		case data := <- dataChan:
			fmt.Printf("data: %d\n", data)
		case err := <- errChan:
			t.Error(err)
		case <- quitChan:
			return
		default:

		}
	}
}