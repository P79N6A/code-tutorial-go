package main

import (
	"net/http"
	"time"
)
/*
client := &http.Client{
CheckRedirect: redirectPolicyFunc,
}

resp, err := client.Get("http://example.com")
// ...

req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...
*/

// http proxy
// custom header
// timeout,connection/read
// follow redirect?
// 并发链接个数, defer加个计数器.....
func main() {
	client := http.Client{}
	resp, err := client.Get("http://example.com")
	req, err := http.NewRequest()

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout, // timeout
	}

}
