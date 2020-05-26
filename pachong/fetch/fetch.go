package fetch

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong StatusCode: %d\n", resp.StatusCode)
	}

	bodyreader := bufio.NewReader(resp.Body)
	e := determinEncoding(bodyreader)

	utf8reader := transform.NewReader(bodyreader, e.NewDecoder())

	return ioutil.ReadAll(utf8reader)

}

func determinEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024) //一般都是1024，自己写的网站短
	if err != nil {
		print(err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e

}
