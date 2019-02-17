package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/knusbaum/globetrotter/globe"
	"net/http"
)

var data = `
PICKLES:
  en.US: Pickles
  de.DE: Gurken
  es.ES: Pepinillos
TOMATO:
  en.US: Tomato
  de.DE: Tomate
  es.ES: Tomate
FRUIT:
  en.US: Fruit
  de.DE: Frucht
  es.ES: Fruta
`

func tryit(g *globe.GlobeDB, str, lang string) {
	fmt.Printf("Looking up %s, %s ... ", str, lang)
	str, err := g.Lookup(str, lang)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Found: %s\n", str)
}


func main() {

	globeDB, err := globe.LoadDB([]byte(data))

	if err != nil {
		fmt.Printf("Failed to load Globetrotter db: %s\n", err)
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/", globe.RequestHandler(globeDB)).Methods("POST")

	http.ListenAndServe("0.0.0.0:8080", r)

}
