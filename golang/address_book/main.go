package main

import (
	pb "example/proto"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io/ioutil"
	"log"
)

const (
	filename = "addressbook.bin"
)

func main() {
	book := &pb.AddressBook{}
	readAddressbook(filename, book)

	fmt.Println("Before: ", book)

	person := &pb.Person{
		Name:  "Vikas Sheoran",
		Id:    1,
		Email: "kkviks@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "123", Type: pb.Person_HOME},
			{Number: "456", Type: pb.Person_MOBILE},
			{Number: "789", Type: pb.Person_WORK},
		},
		LastUpdated: &timestamppb.Timestamp{},
	}

	addPerson(book, person)
	writeAaddressbook(filename, book)

	fmt.Println("After: ", book)
}

func writeAaddressbook(filename string, book *pb.AddressBook) {
	bytes, err := protojson.Marshal(book)
	if err != nil {
		log.Fatalln("Error marshalling address book: ", err)
		return
	}

	if err := ioutil.WriteFile(filename, bytes, 0644); err != nil {
		log.Fatalln("Error writing address book: ", err)
		return
	}
}

func addPerson(book *pb.AddressBook, person *pb.Person) {
	book.People = append(book.People, person)
}

func readAddressbook(filename string, book *pb.AddressBook) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Could not read file: ", filename)
		return
	}

	if err := protojson.Unmarshal(bytes, book); err != nil {
		log.Fatalln("Could not unmarshal file: ", filename)
	}
}
