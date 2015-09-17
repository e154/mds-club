package parser

import (
	"fmt"
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
	total_elements int = 0
	current_element int
	totalChan chan int
	statusChan chan int
	errorChan chan error
	quitChan chan bool
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

func Run(start, stop int) (chan bool, chan int, chan int, chan error) {

	quitChan = make(chan bool, 1)
	totalChan = make(chan int, 1)
	statusChan = make(chan int, 1)
	errorChan = make(chan error, 1)

	total, err := GetTotalElements(URL)
	if err != nil {
		errorChan <- err
		quitChan <- true
	}

	totalChan <- total

	current_element = 0
	go func() {
		scanCatalog(URL, start, stop)
		defer close(quitChan)
		defer close(totalChan)
		defer close(statusChan)
		defer close(errorChan)
		quitChan <- true
	}()

	return quitChan, totalChan, statusChan, errorChan
}