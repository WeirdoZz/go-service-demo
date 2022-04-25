package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func redirectLimitTimes() {
	// 限制重定向次数
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) > 10 {
				return errors.New("重定向次数过多")
			}
			return nil
		},
	}

	request, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/redirect/20", nil)

	client.Do(request)
}

func redirectForbidden() {
	// 禁止重定向
	// 登录请求，防止重定向到首页
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	request, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/cookies/set?name=weirdo", nil)

	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	//content, _ := ioutil.ReadAll(r.Body)
	//fmt.Printf("%s\n", content)

	fmt.Println(r.Request.URL)

	r, _ = client.Do(request)
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", content)
	fmt.Println(r.Request.URL)
}

func main() {
	// 重定向
	//返回一个状态码 3xx
	//redirectLimitTimes()
	redirectForbidden()
}

