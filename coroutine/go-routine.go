package main
import (
	"time"
	"fmt"
)

func teardown() {
        fmt.Println("tear down")
}
func timeoutFun() {
        defer  teardown()
	c1 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
}

func nonBlockFunc() {
	c1 := make(chan string,20)
	go func(c chan string) {
                for {

                        time.Sleep(time.Second )
                        c <- "result 1"
                }
	}(c1)

	for true {
		select {
		case res := <-c1:
			fmt.Println(res)
                        fmt.Println(len(c1))
		// it will not go in here. because time.Second is same.......
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1")
		default:
			fmt.Println("non block")
		}
		time.Sleep(time.Second*2)
	}
}

func main() {
	//timeoutFun()
	nonBlockFunc()
}
