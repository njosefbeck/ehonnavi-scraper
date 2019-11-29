package main

import (
	"fmt"
	"log"

	"github.com/njosefbeck/ehonnavi-scraper/db"
	"github.com/njosefbeck/ehonnavi-scraper/scraperutils"
)

func buildInitialURL(age string, pageNum string) string {
	const first = "https://www.ehonnavi.net/browse_all/list.asp?dp=&st=2&tk="
	return first + age + "&ano=&pg=" + pageNum
}

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	numAdded := 0
	numSkipped := 0
	ages := []string{"00", "01"}

	for _, age := range ages {
		url := buildInitialURL(age, "1")
		scraperutils.ProcessPage(db, age, url, &numAdded, &numSkipped)
	}

	fmt.Println("==========================================")
	fmt.Printf("Added %d new books. Skipped %d books.\n\n", numAdded, numSkipped)
}
