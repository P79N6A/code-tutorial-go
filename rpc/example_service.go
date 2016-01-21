package main

import (
	"fmt"
	pb "code-tutorial-go/rpc/proto"
	"github.com/golang/protobuf/proto"
)
func main() {
	var request pb.GreetRequest
	request.Hobbies = make([]string,  5)
	request.Keyword = make(map[string]int32)
	request.Keyword["e"] = 1
	request.Keyword["f"] = 1
	request.Person = &pb.Person{"alice",12}
	a,err := proto.Marshal(&request)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(a))
	fmt.Println(a)
	fmt.Println("xx")
	fmt.Println(proto.MarshalTextString(&request))
	fmt.Println("xx")

	var re pb.GreetRequest
	proto.Unmarshal(a,&re)
	fmt.Println(re.String())
}
