package main

import "fmt"

func zerovalue(ivalue int) {
	ivalue = 0
}
func zeropointer(ipointer *int) {
	*ipointer = 0
}
func main() {
	i := 1
	fmt.Println("i = ", i)
	zerovalue(i)
	fmt.Println("i zerovalue = ", i)
	zeropointer(&i)
	fmt.Println("i zeropointer value = ", i)
	fmt.Println("i zeropointer address = ", &i)

	msg := "Old Message"
	var msgPointer *string = &msg
	fmt.Println(msgPointer)
	fmt.Println(*msgPointer)
	changeMessage(&msg, "New Message")
	fmt.Println(msg)
}
func changeMessage(mPointer *string, newMessage string) {
	*mPointer = newMessage
}
