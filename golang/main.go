package main

import (
	pb "example/proto"
	"fmt"
)

func main() {
	fmt.Println(doSimple())
	fmt.Println(doComplete())
	fmt.Println(doEnum())
	doOneOf(&pb.Result_Id{Id: 1})
	doOneOf(&pb.Result_Message{Message: "huehuehue"})
}

func doOneOf(m interface{}) {
	switch x := m.(type) {
	case *pb.Result_Id:
		fmt.Println(m.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(m.(*pb.Result_Message).Message)
	default:
		fmt.Printf("%v is neither a Result_Id nor a Message", x)
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_GREEN,
	}
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
