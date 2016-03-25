package main
import (
	"time"
	"net"
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"errors"
	"crypto/tls"
	"net/http/httputil"
)
type Timeout struct {
	Connect   time.Duration
	ReadWrite time.Duration
}
func noRedirect(req *http.Request, via []*http.Request) error {
	return errors.New("No Redirect")
}
// TimeoutDialer
func TimeoutDialer(timeout *Timeout) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, timeout.Connect)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(timeout.ReadWrite))
		return conn, nil
	}
}

func httpsClient() *http.Client {
	to := &Timeout{
		Connect:   time.Duration(1000000) * time.Millisecond,
		ReadWrite: time.Duration(1000000) * time.Millisecond,
	}

	return &http.Client{
		CheckRedirect:noRedirect,
		Transport: &http.Transport{
			Dial: TimeoutDialer(to),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}
// HttpClient
func HttpClient() *http.Client {
	to := &Timeout{
		Connect:   time.Duration(1000000) * time.Millisecond,
		ReadWrite: time.Duration(1000000) * time.Millisecond,
	}

	//http://stackoverflow.com/questions/14661511/setting-up-proxy-for-http-client
	proxyUrl,_ := url.Parse("http://218.106.96.200:80")

	return &http.Client{
		CheckRedirect:noRedirect,
		Timeout:time.Duration(1000000) * time.Millisecond,
		Transport: &http.Transport{
			Dial: TimeoutDialer(to),
			Proxy:  http.ProxyURL(proxyUrl),
			TLSClientConfig:nil,
		},
	}
}

func fetch(_url string, https bool) {
	var client *http.Client
	if https {
		client = httpsClient()
	} else {
		client = HttpClient()
	}
	req,_ := http.NewRequest("GET", _url, nil)


	req.Header.Set("Accept","text/html;q=0.8, */*;q=0.5")
	req.Header.Set("Accept-Charset","utf-8, gbk, gb2312, *;q=0.5")
	req.Header.Set("Accept-Encoding","utf-8")
	req.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
	req.Header.Set("Connection","keep-alive")
	req.Header.Set("Connection","close")
	req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36")
	// dump request....
	dump,_ := httputil.DumpRequest(req, true)
	fmt.Println(string(dump))
	fmt.Println("==========")

	resp,err := client.Do(req)

	dump,_ = httputil.DumpResponse(resp, false)
	fmt.Println(string(dump))
	fmt.Println("==========")

	switch err1 := err.(type) {
	case *url.Error:
		fmt.Println("This is a *url.Error")
		/*
			This is a *url.Error
			Get https://stackoverflow.com/: Service Unavailable
		*/
		fmt.Println(err.Error())
		if err, ok := err1.Err.(net.Error); ok && err.Timeout() {
			fmt.Println("and it was because of a timeout")
		}
	case net.Error:
		if err1.Timeout() {
			fmt.Println("This was a *net.OpError with a Timeout")
		}
	}

	if err != nil && strings.Contains(err.Error(), "use of closed network connection") {
		fmt.Println("Could be from a Transport.CancelRequest")
	}

	if err != nil {
			fmt.Println("XXXXXX")
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header.Get("Location"))
	} else {
		if resp.StatusCode == 200 {
			body, _ := ioutil.ReadAll(resp.Body)
			bodystr := string(body);
			fmt.Println(bodystr)
		}
	}
}
func main() {
//	fetch("http://120.27.109.242/")
//	fetch("https://www.baidu.com/", true)
	//fetch("http://z.cn", false)
	fetch("http://www.baidu.com", true)
}
