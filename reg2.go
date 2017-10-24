package main

import (
        "regexp"
        "fmt"
)

var data string = `<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN"><html><head><title>201 The request has been fulfilled and resulted in a new resource being created</title></head><body><h1>TGT Created</h1><form action="https://passport.dnion.com/v1/tickets/TGT-94233-jiVjefXEa1XmyJw1mvnZ7SQhHDscKXLy0nFkU4cFaZsr35pZIl-passport.dnion.com" method="POST">Service:<input type="text" name="service" value=""><br><input type="submit" value="Submit"></form></body></html>`
func main() {
        r, _ := regexp.Compile("TGT-[0-9]+-[a-zA-Z0-9]+-passport.dnion.com")
        fmt.Println(r.FindString(data))

}
