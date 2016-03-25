package main

type User struct {
Name string
}

func main() {
u := &User{Name: "Leto"}
println(u.Name)
Modify(u)
println(u.Name)
	Modify2(*u)
	println(u.Name)
}

func Modify(u *User) {
u.Name = "Paul"
}
func Modify2(u User) {
	u.Name = "XXX"
}
