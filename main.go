package main

import (
	"fmt"
	"github.com/knusbaum/globetrotter/globe"
)

var data = `
PICKLES:
  en.US: Pickles
  de.DE: Gurken
  es.MX: Pepinillos
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

	tryit(globeDB, "PICKLES", "en.US")
	tryit(globeDB, "PICKLES", "de.DE")
	tryit(globeDB, "PICKLES", "es.MX")

	tryit(globeDB, "PICKLES", "fr.FR")
	tryit(globeDB, "HOT_POCKETS", "en.US")
	
}