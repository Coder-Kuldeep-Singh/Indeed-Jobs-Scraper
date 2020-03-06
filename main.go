//Scraping the All Job Titles from the Indeed.co.in
package main

import (
	"crypto/tls"
	"fmt"
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
	// log.Println(response)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Page response is nil", err)
	}
	// log.Println(string(body))
	return body

}
func GetBrowseJobs(Url string) {
	response, err := http.Get(Url)
	if err != nil {
		log.Println("Error to Connect with Indeed Home page.", err)
	}
	defer response.Body.Close()
	// fmt.Println(response)
	// Load the HTML document
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body", err.Error())
		os.Exit(1)
	}
	document.Find("a.icl-GlobalFooter-link").Each(processElement)
}

func processElement(index int, element *goquery.Selection) {
	//see if the href attribute exists on the element
	href, exists := element.Attr("href")
	if exists {
		// fmt.Println(href)
		BrowseJobsPage(href)
	}
}
func BrowseJobsPage(Urls string) {
	fmt.Println(Urls)
	response, err := http.Get(Urls)
	if err != nil {
		log.Println("Error to Connect with Indeed Browse Jobs Page.", err)
	}
	defer response.Body.Close()
	// fmt.Println(response)
	// Load the HTML document
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body", err.Error())
	}
	document.Find("table#categories tbody tr td a").Each(Processjobs)
	fmt.Println("***********************************************************************")
}

func Processjobs(index int, element *goquery.Selection) {
	//see if the href attribute exists on the element
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
		PerJobsTitlePage(href)
	}
}

func PerJobsTitlePage(Urls string) {
	fmt.Println(Urls)
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: tr}

	// SSL config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := http.Client{Transport: transport}
	response, err := client.Get("https://indeed.co.in" + Urls)
	if err != nil {
		log.Println("Error to Connect with Indeed Jobs Category Page.", err)
	}
	defer response.Body.Close()
	// fmt.Println(response)
	// Load the HTML document
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body", err.Error())
	}
	document.Find("table#titles tbody tr td p.job a").Each(ProcessSinglejob)
	fmt.Println("***********************************************************************")
}

func ProcessSinglejob(index int, element *goquery.Selection) {
	//see if the href attribute exists on the element
	href, exists := element.Attr("title")
	if exists {
		fmt.Println(href)
	}
}

func main() {
	GetBrowseJobs("https://www.indeed.co.in/")
}
