package main

import (
	"fmt"
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

	// find all div in the doc
	doc.Find("div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	// find all div with class sqekv3-5 in the doc
	doc.Find("div.sqekv3-5").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	// find the element in the parent element
	doc.Find("div.sqekv3-3").Each(func(i int, selection *goquery.Selection) {
		// fmt.Println(selection.Text())
		log.Print(selection.Find("div.sqekv3-2").Text()) // category
		log.Print(selection.Find("div.sqekv3-5").Text()) // title
		log.Print(selection.Find("div.sqekv3-6").Text()) // Date
	})

	// get the url (a) element from the dive
	doc.Find("div.sqekv3-3").Each(func(i int, selection *goquery.Selection) {
		a := selection.Find("a")
		qHref, _ := a.Attr("href")
		log.Print(qHref)
	})

	/*
		Some Note:
		selection.Find("div.sqekv3-2") return a select type, need transfer it to sting by .text()
	*/
}
