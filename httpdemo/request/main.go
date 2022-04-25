package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func printBody(r *http.Response) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func requestByParams() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}

	params := make(url.Values)
	params.Add("name", "weirdo")
	params.Add("age", "18")

	request.URL.RawQuery = params.Encode()

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	printBody(resp)

}

func requestByHead() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("user-agent", "chrome")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	printBody(resp)
}

func main() {
	// 如何设置请求的查询参数 http://httpbin.org/get?name=weirdo&age=18
	// 如何定制请求头，比如修改 user-agent
	//requestByParams()
	requestByHead()
}

