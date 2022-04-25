package timeout

import (
	"context"
	"net"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (con net.Conn, e error) {
				return net.DialTimeout(network, addr, 2*time.Second)
			},
			ResponseHeaderTimeout: 5 * time.Second,
			TLSHandshakeTimeout:   2 * time.Second,
			IdleConnTimeout:       60 * time.Second,
		},
	}
	client.Get("http://httpbin.org/delay/10")
}
