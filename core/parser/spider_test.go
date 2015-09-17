package parser

import (
	"testing"
	"time"
)

func TestGetDocument(t *testing.T) {

	var url string = "http://mds-club.ru/cgi-bin/index.cgi?r=84&lang=rus"
	var agent string = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.90 Safari/537.36"
	var timeout time.Duration = 15

	_, err := GetDocument(agent, url, timeout * time.Second)
	if err != nil {
		t.Errorf("error: %s\n", err.Error())
	}
}