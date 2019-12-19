package main

import (
	"errors"
	"errorwrap/module"
	"fmt"
)

func main() {
	fmt.Println("we get error:", module.Aerror)
	fmt.Println("we get sub error:", errors.Unwrap(module.Aerror))
}
