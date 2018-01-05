package main
import (
"fmt"
"time"
)

func main() {
        t,_ := time.Parse("","21-Jul-2016 00:00:00.013")
        fmt.Println(time.Now())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Date())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Second())
	fmt.Println(time.Now().Format(time.RFC822))
	y,m,d := time.Now().Date()
	time.Now().Month()
	fmt.Println(y)
	fmt.Println(int(m))
	fmt.Println(d)
	value := fmt.Sprintf("%s %d, %d 00:00AM",m.String(), d, y)

	// The form must be January 2,2006.
	form := "January 2, 2006 00:00PM"

	// Parse the string according to the form.
	t, _ = time.Parse(form, value)
	fmt.Println(t)

	fmt.Println(t.Unix())
        yestoday := time.Now().AddDate(0,0,-1)
        y,m,d = yestoday.Date()
        fmt.Println(time.Date(y,m,d,0,0,0,0,time.Local).Unix())

        fmt.Println(yestoday.Unix())
        fmt.Println(y)
	fmt.Println(m.String())
	fmt.Println(d)

        fmt.Println(GetCurrentDate())
}


func GetCurrentDate() string {
        // return 20160102
        y,m,d := time.Now().Date()
        h,mi,s := time.Now().Clock()
        return fmt.Sprintf("%d%02d%02d%02d%02d%02d",y,int(m),d,h,mi,s)
}