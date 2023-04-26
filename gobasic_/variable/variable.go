package main

import "fmt"

func main() {
	var name string = "nike"
	age := 15
	weight := 66.9
	height := 173
	pname := "Mr."
	fmt.Println(pname)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(weight)
	fmt.Println(height)
	fmt.Println(pname+name+" age :",age,"weight :",weight,"height :",height)
	fmt.Println(pname+name)
	fmt.Println(age+int(weight))
	fmt.Println(pname+"World")
}
