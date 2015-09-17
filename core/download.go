package core

import (
	"os"
	"fmt"
	"strings"
	"crypto/md5"
	"io"
	"encoding/hex"
	http "net/http"
)

var (
	QuitChan	chan bool
	DataChan	chan int64
	ErrorChan	chan error
)

func checkErr(err error) {
	if err != nil {
		ErrorChan <- err
	}
}

func DownloadHttp(url string) error {

	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]

	h := md5.New()
	io.WriteString(h, fileName)
	s := hex.EncodeToString(h.Sum(nil))

	output, err := os.Create(fmt.Sprintf("/tmp/%s", s))
	if err != nil {
		checkErr(err)
		return err
	}

	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		checkErr(err)
		return err
	}

	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		checkErr(err)
		return err
	}

	DataChan <- n
	QuitChan <- true

	return nil
}

// TODO add ftp dowloader
func DownloadFtp(url string) error {

	return nil
}

func Download(url string) (chan bool, chan int64, chan error) {

	QuitChan = make(chan bool, 1)
	DataChan = make(chan int64, 1)
	ErrorChan = make(chan error, 10)

	if strings.Contains(url, "ftp://") {
		go DownloadFtp(url)
	} else if strings.Contains(url, "http://") {
		go DownloadHttp(url)
	} else {
		ErrorChan <- fmt.Errorf("url sting not valid")
	}

	return QuitChan, DataChan, ErrorChan
}
