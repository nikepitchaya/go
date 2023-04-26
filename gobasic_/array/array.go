package main

import "fmt"

func main() {
	var price [4]int
	var menu []string
	var products [3]string = [3]string{"A", "B", "C"}
	fmt.Println(products[:])
	fmt.Println(price)
	price[0]=1000
	fmt.Println(price)
	menu = append(menu,"200")
	fmt.Println(menu)
}
