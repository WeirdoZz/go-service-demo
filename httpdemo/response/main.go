package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func responseBody(response *http.Response) {
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("%s", content)
}

func status(response *http.Response) {
	fmt.Println(response.StatusCode) // 状态码
	fmt.Println(response.Status)     // 状态描述
}

func header(response *http.Response) {
	// 使用get可以忽略大小写
	fmt.Println(response.Header.Get("content-type"))
}

func encoding(response *http.Response) {
	bufReader := bufio.NewReader(response.Body)
	bytes, _ := bufReader.Peek(1024) // 预取，不会移动reader的真正读取位置
	e, _, _ := charset.DetermineEncoding(bytes, response.Header.Get("content-type"))
	fmt.Println(e)

	// 用编码方式对应的解码方式进行解码
	bodyReader := transform.NewReader(bufReader, e.NewDecoder())
	content, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func main() {
	r, err := http.Get("http://bilibili.com")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	//responseBody(r)
	//status(r)
	//header(r)
	encoding(r)
}
