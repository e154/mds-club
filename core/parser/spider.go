package parser

import (
//	"io/ioutil"
	"golang.org/x/net/html/charset"
	"github.com/PuerkitoBio/goquery"
	http "net/http"
)

func GetDocument(agent, url string) (body *goquery.Document, err error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)

	req.Header.Set("User-Agent", agent)

	resp, err := client.Do(req)
	checkErr(err)

	defer resp.Body.Close()
	checkErr(err)

	// encode
	utf8, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	checkErr(err)

//	res, err := ioutil.ReadAll(utf8)
//	checkErr(err)

	return goquery.NewDocumentFromReader(utf8)
}