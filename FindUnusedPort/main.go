package main

import (
	"log"
	"net"
)

func FindUnusedPort() string {
	//ask kernel free port by :0 port
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		log.Println(err)
		return ""
	}
	listener, err := net.ListenTCP("tcp", addr)
	defer listener.Close()
	if err != nil {
		log.Println(err)
		return ""
	}
	log.Println("Get unused port:", listener.Addr().String())
	return listener.Addr().String()
}
