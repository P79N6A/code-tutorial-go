package main

import (
	"fmt"
)

func main() {
	a := 123
	fmt.Printf("%+10d\n", a)  //+123
	fmt.Printf("%02d\n", a) //+000000123，利用０来补齐位数，而不是空格
        var b  float32 = 123455.223456889065432
        fmt.Printf("%.5f\n",b)
}
