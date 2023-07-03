package main

import (
	"log"

	"github.com/golang/protobuf/proto"
	pb "golang-learning/protobuf/proto"
)

func main() {
	// define the user in protobuf message
	user := &pb.Person{
		Name: "Jimmy",
		Age: 24,
	}
	log.Printf("Person details are: %v", user)

	// marshal the object into protobuf serialized data
	data, err := proto.Marshal(user)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	log.Printf("The marshaled details : %v", data)

	// define another newUser
	newUser := &pb.Person{}
	log.Printf("newUser is: %v", newUser)

	// unmarshal byte array into object
	err = proto.Unmarshal(data, newUser)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	log.Print(newUser.GetName())
	log.Print(newUser.GetAge())
}