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
	ages := []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "19"}

	for _, age := range ages {
		url := buildInitialURL(age, "1")
		scraperutils.ProcessPage(db, age, url, &numAdded, &numSkipped)
	}

	fmt.Println("==========================================")
	fmt.Printf("Total books added: %d. Total books skipped: %d.\n\n", numAdded, numSkipped)
}
