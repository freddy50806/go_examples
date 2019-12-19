package main

import (
	"log"
	"os"
	"path/filepath"

	_ "./package1"
)

func main() {

	dir, err := os.Getwd()
	log.Println("WD Dir : ", dir, " Err : ", err)
	dir2, err := filepath.Abs(filepath.Dir(os.Args[0]))
	log.Println("Dir : ", dir2, " Err : ", err)
}
