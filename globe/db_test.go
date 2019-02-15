package globe

import (
	"os"
	"fmt"
	"io/ioutil"
	"bytes"
	"testing"
)

var data = []byte(`
PICKLES:
  en.US: Pickles
  de.DE: Gurken
  es.MX: Pepinillos
`)

func TestLoadDB(t *testing.T) {
	_, err := LoadDB(data)

	if err != nil {
		t.Errorf("Failed to load globe DB with LoadDB: %s", err)
	}
}

func TestLoadDBFromReader(t *testing.T) {
	reader := bytes.NewReader(data)
	_, err := LoadDBFromReader(reader)

	if err != nil {
		t.Errorf("Failed to load globe DB with LoadDBFromReader: %s", err)
	}
}

func TestLoadDBFromFile(t *testing.T) {
    tmpFile, err := ioutil.TempFile("/tmp", "test.*.yml")
	name := tmpFile.Name()
	
	if err != nil {
		t.Errorf("Failed to open temporary file: %s", err)
	}
	defer os.Remove(name)
	
	_, err = tmpFile.Write(data)
	tmpFile.Close() // Close regardless
	if err != nil {
		t.Errorf("Failed to write temporary file: %s", err)
	}

	_, err = LoadDBFromFile(name)

	if err != nil {
		t.Errorf("Failed to load globe DB with LoadDBFromReader(%s): %s", name, err)
	}
}

func tryit(g *GlobeDB, str, lang string) {
	fmt.Printf("Looking up %s, %s ... ", str, lang)
	str, err := g.Lookup(str, lang)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Found: %s\n", str)
}

func ExampleLookups() {
	globeDB, err := LoadDB(data)

	if err != nil {
		fmt.Printf("Failed to load Globetrotter db: %s\n", err)
		return
	}

	tryit(globeDB, "PICKLES", "en.US")
	tryit(globeDB, "PICKLES", "de.DE")
	tryit(globeDB, "PICKLES", "es.MX")

	tryit(globeDB, "PICKLES", "fr.FR")
	tryit(globeDB, "HOT_POCKETS", "en.US")
	// Output:
	// Looking up PICKLES, en.US ... Found: Pickles
	// Looking up PICKLES, de.DE ... Found: Gurken
	// Looking up PICKLES, es.MX ... Found: Pepinillos
	// Looking up PICKLES, fr.FR ... Error: Translation fr.FR not found for string PICKLES
	// Looking up HOT_POCKETS, en.US ... Error: No such string HOT_POCKETS
}
