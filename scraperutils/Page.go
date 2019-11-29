package scraperutils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func buildSubURL(relURL string) string {
	return "https://www.ehonnavi.net/browse_all/" + relURL
}

func processHTML(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result, err := FromShiftJIS(string(body))
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(result))
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// Book : struct for a book
type Book struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Age   string `json:"age"`
}

// ProcessPage : process ehonnavi page and save book to books map
func ProcessPage(age string, url string, books map[string]Book) {
	fmt.Println("Now processing page: ", url)

	doc, err := processHTML(url)
	if err != nil {
		log.Fatal(err)
	}

	// Each book listing is inside a div with the class 'detailOneCol'
	// Find all of those listings, loop over them and create a Book
	doc.Find(".detailOneCol").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".text h3 a").Text()
		href, _ := s.Find(".text h3 a").Attr("href")

		_, hasKey := books[href]
		if !hasKey {
			books[href] = Book{Title: title, URL: href, Age: age}
		}
	})

	// Call this function again to recursively work our way through
	// each page in the particular age we're already looping through
	doc.Find(".pageSending").First().Find("a").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "次へ→" {
			href, _ := s.Attr("href")
			ProcessPage(age, buildSubURL(href), books)
		}
	})
}
