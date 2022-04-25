package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFile(url, filename string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	n, err := io.Copy(file, r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, err)
}

type Reader struct {
	io.Reader
	Total   int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)

	r.Current += int64(n)
	fmt.Printf("\r进度:%.2f%%", float32(r.Current*100/r.Total))

	return
}

func downloadFileProgress(url, filename string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := &Reader{
		Reader: r.Body,
		Total:  r.ContentLength,
	}

	n, err := io.Copy(file, reader)
	fmt.Println(n, err)
}

func main() {
	//downloadFile("https://weirdozz.github.io/images/weirdo.jpg", "weirdo.jpg")
	downloadFileProgress("https://weirdozz.github.io/images/weirdo.jpg", "weirdo.jpg")
}
