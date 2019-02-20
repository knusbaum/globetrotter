package globe

import (
	"os"
	"fmt"
	"io/ioutil"
	"bytes"
	"testing"
)

var data = []byte(`
version: 1234
translations:
  PICKLES:
    en.US: Pickles
    de.DE: Gurken
    es.ES: Pepinillos
  TOMATO:
    en.US: Tomato
    de.DE: Tomate
    es.ES: Tomate
`)

var badData = []byte(`
PICKLES:
  en.US:
    - Pickles
    - Guac
  - de.DE: Gurken
  - es.MX: Pepinillos
TOMATO:
  - en.US
    - Tomato
  - de.DE
    - Tomate
  - es.ES
    - Tomate
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

func TestLoadDBFail(t *testing.T) {
	_, err := LoadDB(badData)

	if err == nil {
		t.Errorf("Failed to load globe DB with LoadDB: %s", err)
	}
}

func TestLoadDBFromReaderFail(t *testing.T) {
	reader := bytes.NewReader(badData)
	_, err := LoadDBFromReader(reader)

	if err == nil {
		t.Errorf("Expected error from LoadDBFromReader: %s", err)
	}
}

func TestLoadDBFromFileFail(t *testing.T) {
	tmpFile, err := ioutil.TempFile("/tmp", "test.*.yml")
	name := tmpFile.Name()

	if err != nil {
		t.Errorf("Failed to open temporary file: %s", err)
	}
	defer os.Remove(name)

	_, err = tmpFile.Write(badData)
	tmpFile.Close() // Close regardless
	if err != nil {
		t.Errorf("Failed to write temporary file: %s", err)
	}

	_, err = LoadDBFromFile(name)

	if err == nil {
		t.Errorf("Expected error from LoadDBFromReader(%s): %s", name, err)
	}
}

func TestLookupAll(t *testing.T) {
	g, err := LoadDB(data)

	if err != nil {
		t.Errorf("Failed to load globe DB with LoadDB: %s", err)
	}

	//fmt.Printf("Got: %#v\n", g.LookupAll("en.US"))
	expected := map[string]string {
		"PICKLES": "Pickles",
		"TOMATO": "Tomato",
	}

	actual := g.LookupAll("en.US")
	for k, v := range expected {
		if actual[k] != v {
			t.Errorf("Expected actual[%s] == <%s>, but got actual[%s] == <%s>.\n",
				k, v, k, actual[k])
		}
	}

	expected = map[string]string {
		"PICKLES": "Gurken",
		"TOMATO": "Tomate",
	}

	actual = g.LookupAll("de.DE")
	for k, v := range expected {
		if actual[k] != v {
			t.Errorf("Expected actual[%s] == <%s>, but got actual[%s] == <%s>.\n",
				k, v, k, actual[k])
		}
	}
}

func tryLookup(g *GlobeDB, str, lang string) {
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

	tryLookup(globeDB, "PICKLES", "en.US")
	tryLookup(globeDB, "PICKLES", "de.DE")
	tryLookup(globeDB, "PICKLES", "es.MX")

	tryLookup(globeDB, "TOMATO", "en.US")
	tryLookup(globeDB, "TOMATO", "de.DE")
	tryLookup(globeDB, "TOMATO", "es.ES")

	tryLookup(globeDB, "PICKLES", "fr.FR")
	tryLookup(globeDB, "HOT_POCKETS", "en.US")
	// Output:
	// Looking up PICKLES, en.US ... Found: Pickles
	// Looking up PICKLES, de.DE ... Found: Gurken
	// Looking up PICKLES, es.MX ... Error: Translation es.MX not found for string PICKLES
	// Looking up TOMATO, en.US ... Found: Tomato
	// Looking up TOMATO, de.DE ... Found: Tomate
	// Looking up TOMATO, es.ES ... Found: Tomate
	// Looking up PICKLES, fr.FR ... Error: Translation fr.FR not found for string PICKLES
	// Looking up HOT_POCKETS, en.US ... Error: No such string HOT_POCKETS
}
