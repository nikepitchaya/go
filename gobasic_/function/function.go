package main

import "fmt"

func hello() {
	fmt.Println("Hello Nike")
}
//อย่าลืมกำหนดในส่วนของตัวแปรที่ต้องรีเทิรน์ค่าออกมาด้วย
func plus(value1 int, value2 int) int {
	// result := value1+value2
	// fmt.Println(result)
	return value1 + value2
}
func main() {
	result := plus(10, 2)
	fmt.Println(result)
	hello()
}
