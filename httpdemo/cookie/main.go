package main

import (
	"fmt"
	cookiejar2 "github.com/juju/persistent-cookiejar"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

func rrCookie() {
	// 模拟完成一个登录
	// 请求一个页面，传递基本登录信息,将响应的cookie设置到下一次请求中
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// 通过第一次请求获取到cookie
	firstRequest, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/cookies/set?name=weirdo&password=123", nil)
	firstResp, _ := client.Do(firstRequest)
	defer firstResp.Body.Close()
	content, _ := ioutil.ReadAll(firstResp.Body)
	fmt.Printf("%s\n", content)

	// 将第一次请求得到的cookie添加到之后的请求中
	secondRequest, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/cookies", nil)
	for _, cookie := range firstResp.Cookies() {
		secondRequest.AddCookie(cookie)
	}
	secondResp, _ := client.Do(secondRequest)
	defer secondResp.Body.Close()
	content, _ = ioutil.ReadAll(secondResp.Body)
	fmt.Printf("%s\n", content)

}

func jarCookie() {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	resp, _ := client.Get("http://httpbin.org/cookies/set?name=weirdo&password=123")
	defer resp.Body.Close()

	_, _ = io.Copy(os.Stdout, resp.Body)

}

func login(jar http.CookieJar) {
	client := &http.Client{
		Jar: jar,
	}
	resp, _ := client.PostForm("http://localhost:8080/login", url.Values{"username": {"weirdo"}, "password": {"123"}})
	defer resp.Body.Close()

	fmt.Println(resp.Cookies())
	_, _ = io.Copy(os.Stdout, resp.Body)
}

func center(jar http.CookieJar) {
	client := &http.Client{
		Jar: jar,
	}
	resp, _ := client.Get("http://localhost:8080/center")
	defer resp.Body.Close()

	_, _ = io.Copy(os.Stdout, resp.Body)
}

func main() {
	//rrCookie()
	//jarCookie()
	// cookie分类有两种一种会话期cookie 一种是持久性cookie
	//jar, _ := cookiejar.New(nil)
	jar, _ := cookiejar2.New(nil)

	//login(jar)
	center(jar)
	// 这里能做到持久cookie是因为其保存了cookie信息为文件 在用户目录下的 .go-cookie文件中
	_ = jar.Save()
}
