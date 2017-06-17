/*
go run cli/mcdonaldsMenu/main.go
https://godoc.org/golang.org/x/net/html

mongod
mongo
use mcdonalds
db.items.find()

https://github.com/PuerkitoBio/goquery
go get github.com/PuerkitoBio/goquery

cd cli/mcdonaldsMenu && go build
 */

package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"github.com/Zhanat87/go/util"
	"sync"
	"gopkg.in/mgo.v2"
	"github.com/Zhanat87/go/models"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

const (
	URL = "http://mcdonalds.kz"
	DATABASE = "mcdonalds"
	COLLECTION = "items"
)

// Helper function to pull the href attribute from a Token
func getHref(t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return
}

func getMenuLinks(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		panic("ERROR: Failed to crawl \"" + url + "\"")
	}
	b := resp.Body
	defer b.Close() // close Body when the function returns

	var links []string

	z := html.NewTokenizer(b)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return links
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is an <a> tag
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			// Extract the href value, if there is one
			ok, url := getHref(t)
			if !ok {
				continue
			}

			if len(url) > 9 && url[0:9] == "/ru/menu/" && util.InArray(url, links) == false {
				links = append(links, url)
			}
		}
	}
}

func getMcdonaldsItem(url string) *models.McdonaldsItem {
	doc, err := goquery.NewDocument(URL + url)
	if err != nil {
		panic(err)
	}

	price, err := strconv.Atoi(doc.Find("span.price__value").Contents().Text())
	if err != nil {
		panic(err)
	}
	document := &models.McdonaldsItem{
		Title: doc.Find("h1.article__title.article__title--center").Contents().Text(),
		Image: URL + doc.Find("img.article__img.article__img--center").AttrOr("src", ""),
		Price: price,
		Description: doc.Find("div.article__description.article__description--center").Contents().Text(),
	}
	return document
}

func parseItem(url string, wg *sync.WaitGroup, c *mgo.Collection) {
	defer wg.Done()
	c.Insert(getMcdonaldsItem(url))
}

func main() {
	urls := getMenuLinks(URL + "/menu")
	count := len(urls)

	if count > 0 {
		session, err := mgo.Dial("localhost:27017")
		if err != nil {
			panic("ERROR db connect: " + err.Error())
		}
		c := session.DB(DATABASE).C(COLLECTION)
		c.RemoveAll(nil)

		var wg sync.WaitGroup
		wg.Add(count)
		for _, url := range urls {
			go parseItem(url, &wg, c)
		}
		wg.Wait()

		fmt.Println("Done! ", count, " items")
	} else {
		fmt.Println("Menu not found")
	}
}
