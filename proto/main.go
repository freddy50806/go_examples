package main

import (
	"fmt"
	"log"
	pb "proto/pb/golang"

	"github.com/golang/protobuf/proto"
)

func main() {
	metaData, err := proto.Marshal(&pb.AccountInfo{Name: "John", PersonId: 11223344, DeviceName: []string{"Computer", "Phone"}})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(metaData)
	data := pb.AccountInfo{}
	if err := proto.Unmarshal(metaData, &data); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data) //{John 11223344 [Computer Phone] {} [] 0}
}
