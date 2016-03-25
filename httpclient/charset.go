package main
import (
	"net/http"
	"io/ioutil"
	"fmt"
	"io"
	"compress/gzip"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"strings"
	"mustard/base/file"
)
func Gethtml(url string) []byte{
	client := new(http.Client)

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Accept-Encoding", "gzip")

	response, _ := client.Do(request)
	defer response.Body.Close()

	// Check that the server actually sent compressed data
	var reader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		fmt.Println("XXXXXXXXXX gzip")
		reader, _ = gzip.NewReader(response.Body)
		defer reader.Close()
	default:
		reader = response.Body
	}
	if b, err := ioutil.ReadAll(reader); err == nil {
		return b
	}
	return []byte("")
}
func transformString(t transform.Transformer, s string) (string, error) {
	r := transform.NewReader(strings.NewReader(s), t)
	b, err := ioutil.ReadAll(r)
	return string(b), err
}
func main() {
	content := Gethtml("http://www.jb51.net/")
	file.WriteStringToFile(string(content), "./gbk.html")
	//content := Gethtml("http://www.baidu.com")
	e,n,c := charset.DetermineEncoding(content,"text/html")
	println(e)
	println(n)
	println(c)
	if n != "utf-8" {
		if e != nil {
			s, err := transformString(e.NewDecoder(), string(content))
			if err == nil {
				file.WriteStringToFile(s, "./out.html")
			}
		}
	}
}
