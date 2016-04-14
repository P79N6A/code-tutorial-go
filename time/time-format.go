package main
import (
"fmt"
"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Date())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Second())
	fmt.Println(time.Now().Format(time.RFC822))
	y,m,d := time.Now().Date()
	time.Now().Month()
	fmt.Println(y)
	fmt.Println(m.String())
	fmt.Println(d)
	value := fmt.Sprintf("%s %d, %d 00:00AM",m.String(), d, y)

	// The form must be January 2,2006.
	form := "January 2, 2006 00:00PM"

	// Parse the string according to the form.
	t, _ := time.Parse(form, value)
	fmt.Println(t)

	fmt.Println(t.Unix())
}

