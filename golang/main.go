package main

import (
	pb "example/proto"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	//fmt.Println(doSimple())
	//fmt.Println(doComplete())
	//fmt.Println(doEnum())
	//doOneOf(&pb.Result_Id{Id: 1})
	//doOneOf(&pb.Result_Message{Message: "huehuehue"})
	//fmt.Println(doMap())
	fmt.Println(doFile())
}

func doFile() proto.Message {
	filename := "storage.bin"
	message := &pb.EmailPass{Email: "kkviks@gmail.com", Password: "something-secret"}
	WriteToFile(filename, message)
	msg := ReadFromFile(filename, message)
	return msg
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myid":  {Id: 1},
			"myid2": {Id: 2},
			"myid3": {Id: 3},
		},
	}
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
