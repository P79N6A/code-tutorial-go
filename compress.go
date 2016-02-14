package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io/ioutil"
)

func main() {
	txtb,_ := ioutil.ReadFile("/Users/glc/workspace/test/http___bj.58.com_zufang_24725387601214x.shtml")
	txt := string(txtb)
	var buf bytes.Buffer
	bufzip,_ := flate.NewWriter(&buf,9)
	bufzip.Write([]byte(txt))
	bufzip.Flush()
	bufstr := buf.String()
	fmt.Println(len(bufstr))

	newbuf := bytes.NewBuffer([]byte(bufstr))
	zipread := flate.NewReader(newbuf)
	str,_ := ioutil.ReadAll(zipread)
	fmt.Println(len(string(str)))
	fmt.Println(string(str))
}
