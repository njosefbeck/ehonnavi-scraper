package main

import (
	"fmt"
	"log"

	"github.com/njosefbeck/ehonnavi-scraper/scraperutils"
)

func buildInitialURL(age string, pageNum string) string {
	const first = "https://www.ehonnavi.net/browse_all/list.asp?dp=&st=2&tk="
	return first + age + "&ano=&pg=" + pageNum
}

func main() {
	books := map[string]scraperutils.Book{}

	ages := []string{"00", "01"}

	for _, age := range ages {
		url := buildInitialURL(age, "1")
		scraperutils.ProcessPage(age, url, books)
		fmt.Println("Age:", age, "books added. JSON now has", len(books), "items.")
	}

	success, err := scraperutils.WriteToFile(books, "books.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(success)
}
