package main

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
)

func toJSON(pb proto.Message) string {
	bytes, err := protojson.Marshal(pb)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func fromJSON(message string, pb proto.Message) {
	if err := protojson.Unmarshal([]byte(message), pb); err != nil {
		log.Fatalln("Couldn't unmarshal message ", message)
		return
	}

}
