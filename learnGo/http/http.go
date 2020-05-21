package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	/* 基本简单的http请求方法
	resp, err := http.Get("http://www.imooc.com")
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	*/
	request, err2 := http.NewRequest(http.MethodGet, "https://www.imooc.com", nil)
	if err2 != nil {
		panic(err2)
	}
	//装成手机来浏览
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")

	/* 使用默认的Client
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}*/
	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			// client开始工作
			fmt.Printf("Redirect:", req)
			return nil
		},
	}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("the value is %s\n", dump)
}
