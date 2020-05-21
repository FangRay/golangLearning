package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func douban() {
	resp, err := http.Get("http://book.douban.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		println("something wrong")
		println(resp.StatusCode)
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", bytes)

}
