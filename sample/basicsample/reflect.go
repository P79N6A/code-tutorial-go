/*
	reflect
*/
package main
import (
  "fmt"
  "reflect"
)
type TaInterface interface {
    Name() string
    SetName(string)
}
type Ta struct{
  name string
}                                                                                                                         
func (this *Ta)Name() string {
  return this.name
}
func (this *Ta)SetName(name string) {
    this.name = name
}
func main() {
  s := "XXXXX"
  fmt.Println(reflect.TypeOf(s))
  fmt.Println(reflect.ValueOf(s))
  a := &Ta{name:s}
  fmt.Println(reflect.TypeOf(a))
  fmt.Println(reflect.TypeOf(a).NumMethod())
  for m := 0; m < reflect.TypeOf(a).NumMethod(); m++ {
    method := reflect.TypeOf(a).Method(m)
    fmt.Println(method.Type)         // func(*main.MyStruct) string
    fmt.Println(method.Name)         // GetName
    fmt.Println(method.Type.NumIn())
  }
  fmt.Println(reflect.ValueOf(a))
}                                                                                                                         
   
