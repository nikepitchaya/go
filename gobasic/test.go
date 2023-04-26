package main

import "fmt" //แพ็คเกจที่ไว้จัดการการแสดงผล ข้อความ (input/output) หรือจัดการเกี่ยวกับ format มีสถานะเป็น bulitin fucntion
//รูปแบบของภาษา go จะเป็นแบบ static type โดยจะต้องมีการประกาศตัวแปร ก่อนที่จะนำไปใช้งาน

func main() {
	//การกำหนดตัวแปร
	const pname string = "พลทหาร" //นิยามค่าคงที่
	name := "NikeRuksiam"         // นิยามตัวแปรแบบ type inference
	var age int = 25              // นิยามตัวแปรแบบ manual type declaration

	//การแสดงผลโดยใช้ Println
	fmt.Println("My name is", name)
	fmt.Println("My age is", age)

	//การแสดงผลโดยใช้ Printf
	fmt.Printf("My name is %v \n", name)
	fmt.Printf("Datatype is %T \n", name)
	fmt.Printf("My age is %v \n", age)
	fmt.Printf("Datatype is %T \n", age)

	//ตัวดำเนินการทางคณิตศาสตร์
	num1, num2 := 10, 3
	fmt.Println("หาร", num1/num2) // เครื่องหมายอาจจะเป็น + - * / ก็ได้
	fmt.Println("หารเอาเศษ", num1%num2)

	//ตัวดำเนินการเปรียบเทียบ
	fmt.Println("num1 == num2 ? ", num1 == num2) //อาจจะมีตัวดำเนินการอื่นเช่น == != > < >= <=
	fmt.Println("num1 > num2 ?", num1 > num2)
	fmt.Println("num1 != num2 ?", num1 != num2)

	/*รับค่าจากคีย์บอร์ดด้วย Scanf โดยจะมีค่าที่ส่งเข้าไปทำงานอยู่ 2 ค่า คือ (รูปแบบตัวแทนชนิดข้อมูล(string format),ตัวเก็บข้อมูล) */
	// var nickname string
	// fmt.Print("ขอชื่อหน่อยค่ะ : ")
	// fmt.Scanf("%s", &nickname)
	// fmt.Println("สวัสดี", nickname)

	//เงื่อนไข
	//If statement
	// var score int
	// fmt.Print("กรอกคะแนนสอบของมึง : ")
	// fmt.Scanf("%d", &score)
	// if score > 50 {
	// 	fmt.Printf("สกอร์ของคุณคือ %v สอบผ่าน", score)
	// } else if score <= 50 {
	// 	fmt.Printf("สกอร์ของคุณคือ %v สอบตก", score)
	// } else {
	// 	fmt.Print("กรุณากรอกข้อมูลให้ถูกต้อง")
	// }

	//Switch..Case
	// var key int
	// fmt.Print("ป้อนตัวเลขหน่อยครับ : ")
	// fmt.Scanf("%d", &key)
	// switch key {
	// case 1:
	// 	fmt.Print("เติมตังนะครับ")
	// case 2:
	// 	fmt.Print("ถอนตังนะครับ")
	// default:
	// 	fmt.Print("เงื่อนไขไม่ถูกต้องว่ะ")
	// }

	//Array (defual value มีค่าเท่ากับ 0)
	var numbers [3]int = [3]int{100, 200, 300} // ประกาศ Array แบบทั่วไป
	// numberss := [3]int{100,200,300} //ประกาศ Array แบบ Shorthand
	fmt.Println(numbers[0])
	fmt.Println(numbers[2])
	fmt.Println("Size Array : ", len(numbers))

	//Slices (มีความสามารถในการปรับเปลี่ยนขนาดได้ => Dynamic Arrays)
	names := []string{"สมชาย", "แก้วตา"}
	names = append(names, "สมชาติ") // เพิ่มข้อมูลใน Array
	fmt.Println(names[2])
	fmt.Println(names[0:]) //เข้าถึงสมาชิกแบบกำหนดช่วง

	//Map ใช้ key ในการอ้างอิงตำแหน่งของอาร์เรย์แทนเลข index
	meptest := map[string]string{"TH": "ไทย", "EN": "อังกฤษ"}
	fmt.Println(meptest["TH"])
	fmt.Println(meptest["EN"])
	value, check := meptest["EN"]
	if check {
		fmt.Println(value)
	} else {
		fmt.Print("ไม่เจอข้อมูลเด้อสู")
	}

	//For Loop
	for count := 1; count <= 3; count++ {
		fmt.Println("KUY ")
	}
	type items struct {
		name string
		id   int
	}
	testitem := []items{
		{name: "nike", id: 1},
		{name: "aon", id: 2}, {name: "aoddn", id: 3}, {name: "aoasdasn", id: 4}, {name: "aon31", id: 5},
	}
	// testitem = testitem[1:]
	// fmt.Println(testitem)
	var id int
	fmt.Print("Enter Id for Delete : ")
	fmt.Scanf("%d", &id)
	for index, item := range testitem {
		if item.id == id {
			testitem = testitem[index+1:]
		}
		
	}
	fmt.Println(testitem)

}
