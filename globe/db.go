package globe

import (
	"io"
	"os"
	"fmt"
	"errors"
	"gopkg.in/yaml.v2"
)

type GlobeDB struct {
	translations map[string]map[string]string
}


func LoadDBFromFile(filename string) (*GlobeDB, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return LoadDBFromReader(file)
}

func LoadDBFromReader(r io.Reader) (*GlobeDB, error) {
	decoder := yaml.NewDecoder(r)
	g := &GlobeDB{}
	
	err := decoder.Decode(&g.translations)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func LoadDB(in []byte) (*GlobeDB, error) {
	g := &GlobeDB{}
	err := yaml.Unmarshal(in, &g.translations)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *GlobeDB) Lookup(strname, lang string) (string, error) {
	languages, ok := g.translations[strname]
	if !ok {
		return "", errors.New(fmt.Sprintf("No such string %s", strname))
	}

	translation, ok := languages[lang]
	if !ok {
		return "", errors.New(fmt.Sprintf("Translation %s not found for string %s", lang, strname))
	}

	return translation, nil
}


