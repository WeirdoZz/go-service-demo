package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func postForm() {
	// form data 形式 query string 类似于 name=weirdo&age=18
	data := make(url.Values)
	data.Add("name", "weirdo")
	data.Add("age", "18")
	payload := data.Encode()

	r, _ := http.Post("http://httpbin.org/post", "application/x-www-form-urlencoded", strings.NewReader(payload))
	defer r.Body.Close()

	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", content)
}

func postJson() {
	u := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "weirdo",
		Age:  18,
	}
	payload, _ := json.Marshal(u)

	r, _ := http.Post("http://httpbin.org/post", "application/json", bytes.NewReader(payload))
	defer r.Body.Close()

	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", content)
}

func postFile() {
	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)
	_ = writer.WriteField("word", "weirdo")

	upload1Writer, _ := writer.CreateFormFile("uploadfile1", "file1.txt")
	uploadFile1, _ := os.Open("file1.txt")
	defer uploadFile1.Close()
	io.Copy(upload1Writer, uploadFile1)

	upload2Writer, _ := writer.CreateFormFile("uploadfile2", "file2.txt")
	uploadFile2, _ := os.Open("file2.txt")
	defer uploadFile1.Close()
	io.Copy(upload2Writer, uploadFile2)

	writer.Close()
	//fmt.Println(body.String())
	fmt.Println(writer.FormDataContentType())

	resp, _ := http.Post("http://httpbin.org/post", writer.FormDataContentType(), body)
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", content)
}

func main() {
	//post form
	//post json
	//post file
	//postForm()
	//postJson()
	postFile()
}
