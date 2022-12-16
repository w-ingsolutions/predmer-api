package main

import (
	"fmt"
	"log"
	"net/http"

	_ "gioui.org/app/permission/storage"
	"github.com/PuerkitoBio/goquery"
	in "github.com/w-ingsolutions/predmer-api/OLD/cfg/ini"

	"predmer-api/OLD/cfg"
)

func main() {
	w := calc.NewWingCal()
	if cfg.Initial {
		fmt.Println("running initial sync")
	}
	in.Init(w.Podesavanja.File)
	//w.APIimportFS(jsonDBpodradovi,jsonDBradovi)

	nbsAPI()

}

func processElement(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
	fmt.Println(element)
}

func nbsAPI() {
	// Make HTTP request
	response, err := http.Get("https://kursna-lista.com/gedzeti/gadget4transparent.php")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find all links and process them with the function
	// defined earlier
	document.Find("td").Each(processElement)
	fmt.Println(document)
}
