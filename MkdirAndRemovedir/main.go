package main

import "os"

func main() {
	os.MkdirAll("a/b/c", os.ModePerm)
	os.RemoveAll("a")
}
