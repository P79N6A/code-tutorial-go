package main

import ("fmt"
	"time"
)
// GOGCTRACE=1 go run gc.go
// GODEBUG=gctrace=1 ./gc
func main() {
	fmt.Println("getting memory")
	tmp := make([]uint32, 100000000)
	for kk, _ := range tmp {
		tmp[kk] = 0
	}
	time.Sleep(5 * time.Second)
	fmt.Println("returning memory")
	tmp = make([]uint32, 1)
	tmp = nil
	time.Sleep(5 * time.Second)
	fmt.Println("getting memory")
	tmp = make([]uint32, 100000000)
	for kk, _ := range tmp {
		tmp[kk] = 0
	}
	time.Sleep(5 * time.Second)
	fmt.Println("returning memory")
	tmp = make([]uint32, 1)
	tmp = nil
	time.Sleep(5 * time.Second)
	return
}
