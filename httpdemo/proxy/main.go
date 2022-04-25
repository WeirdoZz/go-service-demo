package proxy

import (
	"net/http"
	"net/url"
)

func main() {
	proxyUrl, _ := url.Parse("sock5://127.0.0.1:1080")

	t := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
}
