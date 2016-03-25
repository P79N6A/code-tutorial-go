package main

import (
	"mustard/base/hash"
	"fmt"
)
func ShuffleInt(len int) []byte {
	sli := []byte{}
	for i := 0;i < len; i++ {
		sli = append(sli, byte(i))
	}
	hash.Shuffle(&sli)
	return sli
}
func main() {

	a := []string{"x","y","z","a","b","c"}
	ll := ShuffleInt(len(a))
	for _,i := range ll {
		fmt.Println(a[i])
	}
}
