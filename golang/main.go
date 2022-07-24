package main

import (
	pb "example/proto"
	"fmt"
)

func main() {
	//fmt.Println(doSimple())
	fmt.Println(doComplete())
}

func doComplete() *pb.Complex {
	return &pb.Complex{
		OneFoo: &pb.Foo{Id: 1, Name: "Vikas"},
		ManyFoo: []*pb.Foo{
			{Id: 2, Name: "X"},
			{Id: 3, Name: "Y"},
		},
	}
}

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          1,
		Name:        "Vikas",
		IsSimple:    true,
		SampleLists: []uint32{1, 2, 3},
	}
}
