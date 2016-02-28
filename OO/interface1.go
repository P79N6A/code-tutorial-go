package main
type Item interface {
	GetId() string
	SetId(string)
}
type MyItem struct {
	id string
}
func (i MyItem)GetId() string {
	return i.id
}
func (i MyItem)SetId(s string) {
	i.id = s
}

type YourItem struct {
	id string
}
func (i *YourItem)GetId() string {
	return i.id
}
func (i *YourItem)SetId(s string) {
	i.id = s
}

func main() {
	my := MyItem{"id1"}
	my.SetId("xxx")
	println(my.GetId())
	your := YourItem{"id2"}
	your.SetId("YYY")
	println(your.id)
}
