package main
import "time"
import "fmt"

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	for true {
		select {
		case res := <-c1:
			fmt.Println(res)
			goto END
		// it will not go in here. because time.Second is same.......
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1")
		default:
			fmt.Println("non block")
		}
		time.Sleep(time.Millisecond*100)
	}
	END:
}
