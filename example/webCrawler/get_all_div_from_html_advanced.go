package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
//DEBUG = 0 // enable debug mode
)

func main() {

	url := "https://www.theonion.com/search?blogId=1636079510&q=autopilot"
	crawlWebsite(url)
}

func crawlWebsite(url string) {

	// get the content of the web page
	log.Print("Reading content from url:", url)
	resp, err := http.Get(url) // resp.body = io.Reader
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Done, Processing the content...")
	}
	defer resp.Body.Close() // close the http request when ending the program.

	doc, err := goquery.NewDocumentFromReader(resp.Body) // func NewDocumentFromReader(r io.Reader) (*Document, error)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.sqekv3-3").Each(func(i int, selection *goquery.Selection) {
		// fmt.Println(selection.Text())
		//log.Print(selection.Find("div.sqekv3-2").Text())   // category
		//log.Print(selection.Find("div.sqekv3-5").Text())   // title
		//log.Print(selection.Find("div.sqekv3-6").Text())   // Date
		a := selection.Find("a")
		qHref, _ := a.Attr("href")
		log.Print(qHref)
	})
}
