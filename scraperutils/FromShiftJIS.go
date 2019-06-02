package scraperutils

import (
  "io"
  "io/ioutil"
  "strings"
  "golang.org/x/text/encoding/japanese"
  "golang.org/x/text/transform"
)

// Taken from:
// https://gist.github.com/hyamamoto/db03c03fd624881d4b84s
func transformEncoding( rawReader io.Reader, trans transform.Transformer) (string, error) {
    ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
    if err == nil {
        return string(ret), nil
    } else {
        return "", err
    }
}

// Convert a string encoding from ShiftJIS to UTF-8
func FromShiftJIS(str string) (string, error) {
    return transformEncoding(strings.NewReader(str), japanese.ShiftJIS.NewDecoder())
}