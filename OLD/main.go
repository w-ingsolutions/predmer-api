package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	_ "gioui.org/app/permission/storage"
	"gioui.org/io/system"
	"gioui.org/layout"
	"github.com/PuerkitoBio/goquery"
	"predmer-api/OLD/app/helper"

	calc "predmer-api/OLD/app"
	"predmer-api/OLD/cfg"
	in "predmer-api/OLD/cfg/ini"
)

//go:embed jsondb/radovi/*
var jsonDBradovi embed.FS

//go:embed jsondb/podradovi/*
var jsonDBpodradovi embed.FS

func main() {
	w := calc.NewWingCal()
	if cfg.Initial {
		fmt.Println("running initial sync")
	}
	in.Init(w.Podesavanja.File)
	//w.APIimportFS(jsonDBpodradovi,jsonDBradovi)
	w.APIimportFS(jsonDBradovi, jsonDBpodradovi)

	//fmt.Println("w.Radovi: ", w.Radovi.PodvrsteRadova["26"])

	//fmt.Println(" PodvrsteRadovaPodvrsteRadova::::::::::: ", w.Radovi.PodvrsteRadova)

	// w.APIpozivIzbornik(w.Radovi.PodvrsteRadova)
	// w.GenerisanjeLinkova(w.IzbornikRadova)
	nbsAPI()

	// go func() {
	// 	defer os.Exit(0)
	// 	if err := loop(w); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()
	// app.Main()
}

func loop(w *calc.WingCal) error {
	for {
		select {
		case e := <-w.UI.Window.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				w.UI.Context = layout.NewContext(&w.UI.Ops, e)
				helper.Fill(w.UI.Context, helper.HexARGB(w.UI.Tema.
					Colors["Light"]), w.UI.Context.Constraints.Max)

				//if !w.API.OK {
				//	w.GreskaEkran()
				//} else {
				w.GlavniEkran()
				//}

				e.Frame(w.UI.Context.Ops)
			}
			w.UI.Window.Invalidate()
		}
	}
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
