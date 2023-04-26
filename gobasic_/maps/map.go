package main

import "fmt"

func main() {
	products := map[string]int{"Ipad":1000}
	products["Iphone"] = 5000
	fmt.Println(products)

	//delete
	delete(products,"Iphone")
	fmt.Println(products)

	value := products["Ipad"]
	fmt.Println(value)
	
	//map ซ้อน map
	var courses = make(map[string]map[string]int)
	courses["Android"]=make(map[string]int)
	courses["Android"]["price"]=200
	courses["Android"]["code"]=888
	fmt.Println(courses)
	fmt.Println(courses["Android"])
}
