package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
    "io"
    "os"
    "compress/gzip"
  "time"
  "net"
  "net/url"
)
// not follow redirect....
func Gethtml3(url string) {
  client := new(http.Client)

  request, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
  request.Header.Add("Accept-Encoding", "gzip")

  response, _ := client.Do(request)
  defer response.Body.Close()
  for k,v := range response.Header {
    fmt.Println(k)
    fmt.Println(v)
  }

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

  io.Copy(os.Stdout, reader)
}
func GetHTML(url string) {
  res, err := http.Get(url)
  if err != nil {
    return
  }
  body, _ := ioutil.ReadAll(res.Body)//转换byte数组
  defer res.Body.Close()
  //io.Copy(os.Stdout, res.Body)//写到输出流，
  bodystr := string(body)//转换字符串
  fmt.Println(bodystr)
}
func GetHtml2(url string) {
  client := &http.Client{}
  reqest, _ := http.NewRequest("GET", url,nil)

  reqest.Header.Set("Accept","text/html;q=0.8, */*;q=0.5")
  reqest.Header.Set("Accept-Charset","utf-8, gbk, gb2312, *;q=0.5")
  reqest.Header.Set("Accept-Encoding","utf-8")
  reqest.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
  reqest.Header.Set("Connection","keep-alive")

  response,_ := client.Do(reqest)
  for k,v := range response.Header {
    fmt.Println(k)
    fmt.Println(v)
  }
  if response.StatusCode == 200 {
    body, _ := ioutil.ReadAll(response.Body)
    bodystr := string(body);
    fmt.Println(bodystr)
  }
}
func GetHtmlTimeout(url string) {
  to := time.Duration(time.Millisecond)
  client := http.Client{
    Timeout:to,
  }
  req,_ := http.NewRequest("GET", url, nil)
  resp,err := client.Do(req)
  if err != nil {
      fmt.Println("+++" + err.Error())
  } else {
    if resp.StatusCode == 200 {
      body, _ := ioutil.ReadAll(resp.Body)
      bodystr := string(body);
      fmt.Println(bodystr)
    }
  }
}

func main() {
  GetHtmlTimeout("http://stackoverflow.com/")
}