package main
type Item struct {
    name string
}
type Queue struct {
    nodes []*Item
}
func main() {
    var q = Queue{}
    println(len(q.nodes))
}
