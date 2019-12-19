package pack

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	dir, err := os.Getwd()
	log.Println("pack1 Dir : ", dir, " Err : ", err)
	dir2, err := filepath.Abs(filepath.Dir(os.Args[0]))
	log.Println("pack1 Dir : ", dir2, " Err : ", err)
}
