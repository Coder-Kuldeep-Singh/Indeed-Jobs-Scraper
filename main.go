//Scraping the All Job Titles from the Indeed.co.in
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"

)

//Fetch function reads the domain and return the response of the page as  bytes
func Fetch(Url string) []byte {
	response, err := http.Get(Url)
	if err != nil {
		log.Println("Error to Connect with Indeed.", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Page response is nil", err)
	}
	// log.Println(string(body))
	return body

}
func GetBrowseJobs(Url string) {
	response := Fetch(Url)
	// Load the HTML document
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Fatal("Error loading HTTP response body", err.Error())
		os.Exit(1)
	}
	document.Find("a.icl-GlobalFooter-link").Each(processElement)
	// log.Println(string(response))
}

func processElement(index int, element *goquery.Selection) {
	//see if the href attribute exists on the element
	href, exists := element.Attr("href")
	if exists {
		log.Println(href)
	}
}

func main() {
	GetBrowseJobs("https://www.indeed.co.in/")
}
