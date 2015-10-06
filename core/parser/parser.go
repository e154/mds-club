package parser

import (
	"fmt"
	"github.com/e154/console"
	"github.com/PuerkitoBio/goquery"
	"../models"
	"time"
	"strings"
	"strconv"
)

const (
	URL = "http://mds-club.ru/cgi-bin/index.cgi?r=84&lang=rus"
	AGENT = "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
	SLEEP = 200			// millisecond
	TIMEOUT = 15		// seconds
)

var (
	total_elements int = 1312
	current_element int
	totalChan chan int
	statusChan chan int
	errorChan chan error
	quitParserChan chan bool
	status string = "stopped"//launched|stopped
)

type Element struct {
	Url			string
	Author		string
	Book		string
	Date		string
	Station		string
	Files 		[]*models.File
}

func checkErr(err error) {
	if err != nil {
		errorChan <- err
		return
	}
}

func getDocument(url string) (*goquery.Document, error) {
	return GetDocument(AGENT, url, time.Duration(TIMEOUT * time.Second))
}

func getNextUrl(doc *goquery.Document) (string, bool) {
	return doc.Find("#main #body_content #roller_active").First().NextFiltered("#roller_passive").Find("a").Attr("href")
}

func getUrl(doc *goquery.Selection) (string, bool) {
	return doc.Find("a").Attr("href")
}

func getElementsFromPage(doc *goquery.Document) ([]*Element, error) {

	elements := make([]*Element, 0)

	table := doc.Find("#catalogtable").Find("tbody").First()
	trs := table.Find("tr.w")
	for tr_i := range trs.Nodes {

		tr := trs.Eq(tr_i)
		node := tr.Find("td")

		element := new(Element)
		url, b := getUrl(node.Eq(0))
		if b {
			element.Url = url
		}

		element.Author = node.Eq(1).Text()
		element.Book = node.Eq(2).Text()
		element.Date = node.Eq(3).Text()
		element.Station = node.Eq(5).Text()
		elements = append(elements, element)
	}

	return elements, nil
}

func getFiles(doc *goquery.Document) ([]*models.File, error) {

	files := make([]*models.File, 0)

	table := doc.Find("#catalogtable").Find("tbody").First()
	trs := table.Find("tr.w")
	for tr_i := range trs.Nodes {
		tr := trs.Eq(tr_i)
		node := tr.Find("td")

		file := new(models.File)

		url, b := getUrl(node.Eq(3))
		if b {
			file.Url = url
		}

		file.Name = node.Eq(3).Text()
		file.Size = node.Eq(4).Text()

		files = append(files, file)
	}

	return files, nil
}

func parseDate(val string) (time.Time, error) {

	valArr := strings.Split(val, ".")
	day, err := strconv.Atoi(valArr[0])
	if err != nil {
		return time.Now(), err
	}

	month, err := strconv.Atoi(valArr[1])
	if err != nil {
		return time.Now(), err
	}

	year, err := strconv.Atoi(valArr[2])
	if err != nil {
		return time.Now(), err
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}

func scanCatalog(url string, start, stop int) error {

	page, err := getDocument(url)
	if err != nil {
		checkErr(err)
		return err
	}

	elements, err := getElementsFromPage(page)
	if err != nil {
		checkErr(err)
		return err
	}

	for _, element := range elements {

		current_element += 1
		statusChan <- current_element

		if start != 0 && current_element < start {
			time.Sleep(time.Millisecond * 100)
			continue
		}

		if element.Url != "" {
			element_page, err := getDocument(element.Url)
			if err != nil {
				checkErr(err)
				continue
			}

			element.Files, err = getFiles(element_page)
			if err != nil {
				checkErr(err)
				continue
			}
		}

		// save author
		// ----------------------------------------------------
		author, err := models.AuthorGet(element.Author)
		if err != nil || author == nil {
			author = new(models.Author)
			author.Name = element.Author
			author.Low_name = strings.ToLower(element.Author)
			author.Save()
		}

		if author.Id == 0 {
			err = fmt.Errorf("author not init: %s\n", author.Name)
			checkErr(err)
//			return err
			continue
		}

		// station
		// ----------------------------------------------------
		station, err := models.StationGet(element.Station)
		if err != nil || station == nil {
			station = new(models.Station)
			station.Name = element.Station
			station.Save()
		}

		if station.Id == 0 {
			err = fmt.Errorf("station not init: %s\n", station.Name)
			checkErr(err)
//			return err
			continue
		}

		// book
		// ----------------------------------------------------
		book, err := models.BookGet(element.Book)
		if err != nil || book == nil {

			book = new(models.Book)
			book.Name = element.Book
			book.Station_id = station.Id
			book.Author_id = author.Id
			book.Url = element.Url
			book.Low_name = strings.ToLower(element.Book)

			date, err := parseDate(element.Date)
			if err == nil {
				book.Date = date
			}

			book.Save()

		} else {

			if assigned := author.IsAssigned(book); !assigned {
				author.AddBook(book)
			}
		}

		if book.Author_id == 0 {
			err = fmt.Errorf("book not assigned to author: %s\n", author.Name)
			checkErr(err)
			return err
		}

		if book.Id == 0 {
			err = fmt.Errorf("book not init: %s\n", book.Name)
			checkErr(err)
			return err
		}

		// save/remove files
		// ----------------------------------------------------
		old_files, err := book.Files()
		if err != nil {
			checkErr(err)
			return err
		}

		if len(old_files) == 0 {
			for _, file := range element.Files {
				if _, err := file.Save(); err != nil {
					checkErr(err)
				} else {
					book.AddFile(file)
				}
			}
		} else {

			// ----------------------------------
			for _, file := range old_files {

				var exist bool
				for _, file2 := range element.Files {
					if file.Name == file2.Name {
						exist = true
						break
					}
				}

				if !exist {
					if err := file.Remove(); err != nil {
						checkErr(err)
						file = nil
					}
				}
			}

			for _, file := range element.Files {

				var exist bool
				for _, file2 := range old_files {
					if file.Name == file2.Name {
						exist = true
						break
					}
				}

				if !exist {
					if _, err := file.Save(); err != nil {
						checkErr(err)
					} else {
						book.AddFile(file)
					}
				}
			}
			// ----------------------------------
		}

		if stop != 0 && current_element >= stop {
			return nil
		}

		select {
		case <-quitParserChan:
			return fmt.Errorf("interrupt")
		default:
		}

		time.Sleep( time.Duration(SLEEP) * time.Millisecond)
	}

	next_url, b := getNextUrl(page)
	if b {
		if err := scanCatalog(next_url, start, stop); err != nil {
			return err
		}
	}

	return nil
}

func GetTotalElements(url string) (int, error) {

	if total_elements != 0 {
		return total_elements, nil
	}

	var total int

	for {
		page, err := getDocument(url)
		if err != nil {
			checkErr(err)
			return 0, err
		}

		elements, err := getElementsFromPage(page)
		if err != nil {
			checkErr(err)
			return 0, err
		}

		total += len(elements)

		next_url, b := getNextUrl(page)
		if !b {
			break
		}
		url = next_url
	}

	total_elements = total

	return total, nil
}

func init() {
	quitParserChan = make(chan bool, 99)
	statusChan = make(chan int)
	totalChan = make(chan int)
	errorChan = make(chan error)

	c_ptr := console.GetPtr()
	c_ptr.AddCommand("scan", func(key, value string, help *string){

		if value == "help" {*help = "start scanning http://mds-club.ru, very slow process"}
		usage_start := "usage: scan <start|stop:string> <start:int> <stop:int>"

		cmd := strings.Split(value, " ")
		if (cmd[0] == "start" && len(cmd) != 3) || (cmd[0] == "stop" && len(cmd) != 1) {
			c_ptr.Printf(usage_start)
			return
		}

		info := func(s string) {
			c_ptr.Printf("{\"scan_status\":{\"info\":\"%s\"}}", s)
		}

		launched := func() {
			status = "launched"
			c_ptr.Printf("{\"scan_status\":{\"status\":\"launched\"}}")
		}

		stopped := func() {
			status = "stopped"
			c_ptr.Printf("{\"scan_status\":{\"status\":\"stopped\"}}")
		}

		switch cmd[0] {
		case "start":

			start, err := strconv.Atoi(cmd[1])
			stop, err := strconv.Atoi(cmd[2])
			if err != nil {
				c_ptr.Printf(usage_start)
				return
			}

			if status == "launched" { return }
			launched()

			current_element = 0
			go func() {

				info("get the number of records")
				total, err := GetTotalElements(URL)
				if err != nil {
					errorChan <- err
					return
				}

				totalChan <- total
				info(fmt.Sprintf("total records: %d", total))

				// scan
				err = scanCatalog(URL, start, stop)
				if err != nil {
					errorChan <- err
				}

				stopped()
			}()

		case "stop":
			if status != "launched" { return }
			quitParserChan <- true
		case "status":
			c_ptr.Printf("{\"scan_status\":{\"status\":\"%s\"}}", status)
		default:
		}
	})

	go func() {
		current_element = 0
		var total int
		for  {
			select {
			case current := <- statusChan:
				c_ptr.Printf("{\"scan_status\":{\"total\":%d,\"current\":%d}}", total, current)
			case t := <- totalChan:
				total = t
			case err := <- errorChan:
				c_ptr.Printf("{\"scan_status\":{\"error\":\"%s\"}}", err.Error())
			default:
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
}