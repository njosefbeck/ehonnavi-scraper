package scraperutils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	database "github.com/njosefbeck/ehonnavi-scraper/db"
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

// ProcessPage : process ehonnavi page and save book to books map
func ProcessPage(db *gorm.DB, age string, url string, totalNumAdded *int, totalNumSkipped *int) {
	fmt.Println("Processing: ", url)

	doc, err := processHTML(url)
	if err != nil {
		log.Fatal(err)
	}

	numAdded := 0
	numSkipped := 0

	// Each book listing is inside a div with the class 'detailOneCol'
	// Find all of those listings, loop over them and create a Book
	doc.Find(".detailOneCol").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".text h3 a").Text()
		url, _ := s.Find(".text h3 a").Attr("href")

		var book = database.Book{}
		// Check to see if book already exists in db (using url)
		// If it exists, FirstOrCreate returns the existing record
		// If it DOESN'T exist, use Attrs() to specify the values to be saved to the db,
		// then FirstOrCreate creates the record
		// Note: IsNew is a virtual field that's not saved to the db
		db.
			Where(database.Book{URL: url}).
			Attrs(database.Book{Title: title, URL: url, Age: age, IsNew: true}).
			FirstOrCreate(&book)

		// Use IsNew field to determine if the book
		// is newly created or already exists
		if book.IsNew {
			numAdded++
		} else {
			numSkipped++
		}
	})

	*totalNumAdded = *totalNumAdded + numAdded
	*totalNumSkipped = *totalNumSkipped + numSkipped

	fmt.Printf("→ Added: Page (%d) Total (%d). Skipped: Page (%d) Total (%d).\n\n", numAdded, *totalNumAdded, numSkipped, *totalNumSkipped)

	// Call this function again to recursively work our way through
	// each page in the particular age we're already looping through
	doc.Find(".pageSending").First().Find("a").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "次へ→" {
			href, _ := s.Attr("href")
			ProcessPage(db, age, buildSubURL(href), totalNumAdded, totalNumSkipped)
		}
	})

}
