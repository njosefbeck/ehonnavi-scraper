package scraperutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Book : struct for a book
type Book struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Age   string `json:"age"`
}

// WriteToFile : write books in memory as JSON to file
func WriteToFile(books map[string]Book, fileName string) (string, error) {
	data, err := json.Marshal(books)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return "", err
	}
	numBooks := fmt.Sprintf("%d", len(books))
	return numBooks + " books saved to " + fileName, nil
}
