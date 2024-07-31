package main

import (
	"log"

	pb "golang-learning/protobuf/proto"

	"google.golang.org/protobuf/proto"
)

func main() {
	// define the user in protobuf message
	user1 := &pb.Person{
		Name: "Jimmy",
		Age:  24,
	}
	log.Printf("Person details are: %v", user1)

	// marshal the object into protobuf serialized data
	data, err := proto.Marshal(user1)
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

	// define another user for compare
	user2 := &pb.Person{
		Name: "Jimmy",
		Age:  24,
	}

	log.Printf("Are user1 and user 2 equal? %v", proto.Equal(user1, user2))

	// person with address details
	detail := make(map[string]string)
	detail["address"] = "26, Rangon Road"
	detail["postalCode"] = "12780"
	detail["state"] = "Penang"
	detail["country"] = "Malaysia"

	user3 := &pb.Person{
		Name: "Tim",
		Age:  37,
		Address: &pb.Details{
			Detail: detail,
		},
	}

	log.Printf("The user3's address is %v", user3.GetAddress())
}
