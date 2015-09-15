package parser

import (
//	"io/ioutil"
	"golang.org/x/net/html/charset"
	"github.com/PuerkitoBio/goquery"
	http "net/http"
	"time"
)

func GetDocument(agent, url string, timeout time.Duration) (body *goquery.Document, err error) {

	client := &http.Client{Timeout:timeout}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		checkErr(err)
		return
	}

	req.Header.Set("User-Agent", agent)

	resp, err := client.Do(req)
	if err != nil {
		checkErr(err)
		return
	}

	defer resp.Body.Close()
	if err != nil {
		checkErr(err)
		return
	}

	// encode
	utf8, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		checkErr(err)
		return
	}

//	res, err := ioutil.ReadAll(utf8)
//	checkErr(err)

	return goquery.NewDocumentFromReader(utf8)
}