package main
import (
	"fmt"
	"strings"
)

func main() {
	s := " asdfe w; "
	fmt.Println(s[:len(s)-1])
	fmt.Println(strings.TrimSpace(s))
}
