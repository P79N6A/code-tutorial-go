package main

import (
	"regexp"
	"fmt"
)

func main() {
	r,_ := regexp.Compile("a([a-z]+)b([a-z])")
	ret := r.FindAllStringSubmatch("abbbzz acbd", -1)
	fmt.Println(ret)

}
