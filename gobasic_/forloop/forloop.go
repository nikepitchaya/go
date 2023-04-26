package main

import (
	"fmt"
)

func main() {
	foLoop()
	// whLoop()
	foEach()
}
func foLoop() {
	//forloop
	var count int = 5
	for i := 0; i < count; i++ {
		fmt.Println("Hello World")
	}
}

// func whLoop() {
// 	//for like whileloop
// 	var i int = 1
// 	for i < 2 {
// 		fmt.Println("KERORO")
// 	}
// }n
func foEach() {
	//for like foreach
	courses := []string{"A", "B", "C", "D", "F","G"}
	for index, item := range courses {
		fmt.Printf("%d. %s \n", index+1, item)
	}
	for _, item := range courses {
		fmt.Printf("%s \n ", item)
	}
}
