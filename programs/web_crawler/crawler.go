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

	url := "https://www.theonion.com/search?blogId=1636079510&q=game%20of%20thrones"
	crawlWebsite(url)

}

func crawlWebsite(url string) {

	// get the content of the web page
	log.Print("Reading content from url:", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Done, Processing the content...")
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// div.sqekv3-3 is the div for each doc
	doc.Find("div.sqekv3-3").Each(func(i int, selection *goquery.Selection) {

		fmt.Println(selection.Find("div.sqekv3-2").Text()) // category
		fmt.Println(selection.Find("div.sqekv3-5").Text()) // title
		fmt.Println(selection.Find("div.sqekv3-6").Text()) // Date
		a := selection.Find("div.sqekv3-5").Find("a")      // find the href link in the div.sqekv3-6
		targetURL, _ := a.Attr("href")
		fmt.Println(targetURL) // the link of the doc
	})
}
