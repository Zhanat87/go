package util

import (
	"encoding/xml"
	"io/ioutil"
	"bytes"
	"log"
	"net/http"
	"encoding/json"
	"golang.org/x/net/html"
)

type H map[string]interface{}

// MarshalXML allows type H to be used with xml.Marshal
func (h H) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{
		Space: "",
		Local: "map",
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for key, value := range h {
		elem := xml.StartElement{
			Name: xml.Name{Space: "", Local: key},
			Attr: []xml.Attr{},
		}
		if err := e.EncodeElement(value, elem); err != nil {
			return err
		}
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func FetchURL(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("unable to GET '%s': %s", url, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("unable to read body '%s': %s", url, err)
	}
	return body
}

func ParseXML(xmlDoc []byte, target interface{}) {
	reader := bytes.NewReader(xmlDoc)
	decoder := xml.NewDecoder(reader)
	if err := decoder.Decode(target); err != nil {
		log.Fatalf("unable to parse XML '%s':\n%s", err, xmlDoc)
	}
}

func InArray(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func ParseJsonData(encoded []byte) map[string]interface{} {
	var j map[string]interface{}
	err := json.Unmarshal(encoded, &j)
	if err != nil {
		panic(err)
	}
	return j
}

func GetHtmlElementAttrValue(t html.Token, attr string) (ok bool, v string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == attr {
			v = a.Val
			ok = true
		}
	}

	return
}
