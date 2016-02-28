package main

import (
	"mustard/base/string_util"
)

func main() {
	s := "xxxxxxxxxx"
	cs,_ := string_util.Compress(s)
	println(len(cs))
	ns,_ := string_util.Uncompress(cs)
	println(ns)
}
