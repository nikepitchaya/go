package main

import "fmt"

func main() {
	// myMoney := 100
	// if myMoney > 100 {
	// 	fmt.Println("Can buy")
	// } else {
	// 	fmt.Println("Can't buy")
	// }
	var money int
	fmt.Print("ใส่เงินสิ")
	fmt.Scanf("%d",&money)
	fmt.Println("KOE",money)
	if money >=100{
		fmt.Println("KUY")
	}else if money<100 && money>=0{
		fmt.Println("KUY YAI")
	}else{
		fmt.Println("KUY lek")
	}
}
