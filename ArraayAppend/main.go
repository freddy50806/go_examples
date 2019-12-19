package main

import "fmt"

var list []string

func main() {
	list = []string{"a", "b", "c", "d", "e"}
	fmt.Println(list[:0])
	fmt.Println(list[:1])
	list = append(list[:3], list[4:]...) ///a,b,c,e
	fmt.Println(list, len(list))
	list = append(list[:0], list[1:]...) ///b,c,e
	fmt.Println(list, len(list))

	list = append(list[:2], list[3:]...)
	fmt.Println(list, len(list), cap(list))

}
