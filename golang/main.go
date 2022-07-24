package main

import "fmt"
import pb "example/proto"

func main() {
	fmt.Println(doSimple())
}

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          1,
		Name:        "Vikas",
		IsSimple:    true,
		SampleLists: []uint32{1, 2, 3},
	}
}
