package main

import (
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func WriteToFile(filename string, pb proto.Message) {
	bytes, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Couldn't marshal message: ", pb)
	}

	if err := ioutil.WriteFile(filename, bytes, 0644); err != nil {
		log.Fatalln("Error writing message to file: ", err)
		return
	}

	//fmt.Printf("Message %v successfully written to file %v", pb, filename)
}

func ReadFromFile(filename string, pb proto.Message) proto.Message {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading message from file: ", err)
		return nil
	}

	if err := proto.Unmarshal(file, pb); err != nil {
		log.Fatalln("Error in unmarshal message from file ", filename)
		return nil
	}

	return pb
}
