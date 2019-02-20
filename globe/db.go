package globe

import (
	"io"
	"os"
	"fmt"
	"errors"
	"gopkg.in/yaml.v2"
)

type GlobeDB struct {
	Version int
	Translations map[string]map[string]string
}

type TranslationPair struct {
	Key string
	String string
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

	err := decoder.Decode(&g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func LoadDB(in []byte) (*GlobeDB, error) {
	g := &GlobeDB{}
	err := yaml.Unmarshal(in, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *GlobeDB) Lookup(strname, lang string) (string, error) {
	languages, ok := g.Translations[strname]
	if !ok {
		return "", errors.New(fmt.Sprintf("No such string %s", strname))
	}

	translation, ok := languages[lang]
	if !ok {
		return "", errors.New(fmt.Sprintf("Translation %s not found for string %s", lang, strname))
	}

	return translation, nil
}


// Should optimize this. Right now it's O(Num_langs + Num_Keys)
// Basically a full scan is necessary to assemble the dictionary.
func (g *GlobeDB) LookupAll(lang string) map[string]string {
	dictionary := make(map[string]string)

	for key, strMap := range g.Translations {
		translation, ok := strMap[lang]
		if ok {
			//pairs = append(pairs, TranslationPair{key, translation})
			dictionary[key] = translation
		}
	}
	return dictionary
}
