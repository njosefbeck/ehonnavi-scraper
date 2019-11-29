package scraperutils

import (
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"strings"
)

// Taken from:
// https://gist.github.com/hyamamoto/db03c03fd624881d4b84s
func transformEncoding(rawReader io.Reader, trans transform.Transformer) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
	if err == nil {
		return string(ret), nil
	}

	return "", err
}

// FromShiftJIS : Convert a string encoding from ShiftJIS to UTF-8
func FromShiftJIS(str string) (string, error) {
	return transformEncoding(strings.NewReader(str), japanese.ShiftJIS.NewDecoder())
}
