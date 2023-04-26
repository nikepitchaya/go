package main

import (
	// "database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb" //ต้องใส่ Blank Unditifier ด้วยเพื่อให้ go sql ใช้
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Test struct {
	pname string
	fname string
	lname string
	age   int
}

var db *sqlx.DB

func main() {
	fmt.Println("My Drivers is MySQL")
	var err error // ประกาศเพื่อให้ db ข้างในเป็นตัวเดียวกับข้างนอก
	db, err = sqlx.Open("mysql", "root:@tcp(localhost:3306)/gosql")
	if err != nil {
		panic(err)
	}

	//Add Data
	// test := Test{"Dr.", "K", "AM", 120}
	// err = AddTest(test)
	// if err != nil {
	// 	panic(err)
	// }

	//Update Data
	// uptest := Test{"Mr.", "AM", "PM", 120}
	// err = UpdateTest(uptest)
	// if err != nil {
	// 	panic(err)
	// }

	//Delete Data
	// err = DeleteTest(50)
	// if err != nil {
	// 	panic(err)
	// }

	//Get All Data by sqlx
	tests, err := GetTestsX()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("Data is : \n", tests)

	//Get All Data
	// tests, err := GetTests()
	// if err != nil {
	// 	panic(err)
	// 	return
	// }
	// fmt.Println("Data is : \n", tests)

	//Get One Row Data
	// tests, err := GetTest(28)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(tests)

	// var (
	// 	pname string
	// 	fname string
	// 	lname string
	// 	age   int
	// )
	// err = db.QueryRow("SELECT * from test where age=28 ").Scan(&pname, &fname, &lname, &age)
	// if err != nil {
	// 	log.Fatal("Unable", err)
	// }
	// fmt.Printf("ชื่อ : %s %s %s อายุ : %d \n", pname, fname, lname, age)

	// test := []struct {
	// 	pname string
	// 	lname string
	// 	fname string
	// 	age   int
	// }{
	// 	{"Mr.", "Hae", "Sang", 24},
	// 	{"Mr.", "Na", "Mee", 50},
	// 	{"Mr.", "Kie", "Jae", 40},
	// }
	// stmt, err := db.Prepare("INSERT INTO test VALUE(?,?,?,?)")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// for _, test := range test {
	// 	_, err = stmt.Exec(test.pname, test.lname, test.fname, test.age)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

}

//ดึงข้อมูลโดยใช้ sqlx
func GetTestsX() ([]Test, error) {
	
	query := "select * from test"
	test := []Test{}
	err := db.Select(&test, query) //รับค่าเป็น object ปลายทางที่เราต้องการที่จะได้ค่ามันแค่ส่งตำแหน่งไป
	if err != nil {
		return nil, err
	}
	return test, nil
}

//ดึงข้อมูลออกมา
func GetTests() ([]Test, error) {
	err := db.Ping() // เทสว่าดาต้าเบสทำงานอยู่ไหม
	if err != nil {
		return nil, err //ผู้ที่เรียกใช้ฟังก์ชันควรจะ handle err ด้วยตัวเอง
	}
	defer db.Close() //ปิดการเชื่อมต่อ
	fmt.Println("Success Database Connected")

	result, err := db.Query("select * from test") //การประกาศตัวแปรหาก ตัวแปรถูกประกาศไปแล้วจะต้องใช้ = แต่หากยังไม่ถูกประกาศสามารถใช้ := ได้ แต่หากมีตัวใดตัวนึงที่ยังไม่ถูกประกาศก็จะสามารถใช้ := ได้เช่นกัน
	if err != nil {
		return nil, err
	}
	tests := []Test{}
	for result.Next() {
		// pname := ""
		// lname := ""
		// fname := ""
		// age := 0
		testt := Test{}
		err = result.Scan(&testt.pname, &testt.fname, &testt.lname, &testt.age)
		if err != nil {
			return nil, err
		}
		tests = append(tests, testt)

	}
	return tests, nil
}

//ดึงข้อมูลออมา 1 row โดย sqlx
func GetTestX(age int) (*Test, error) {
	query := "select * from  test where age=?"
	test := Test{}
	err := db.Get(&test, query, age)
	if err != nil {
		return nil, err
	}
	return &test, nil
}

//ดึงข้อมูลออกมา 1 row
func GetTest(age int) (*Test, error) { //
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	query := "select * from test where age=?" //argument จะเรียงตาม index เพราะฉะนั้นตอน Query ข้อมูลจำเป็นต้องให้ตำแหน่งของข้อมูลเรียงตาม index
	row := db.QueryRow(query, age)            //ใช้สำหรับการดึงข้อมูลมาแค่ 1 แถวโดยจะรีเทริน์แค่แถวๆนั้นมา
	test := Test{}
	err = row.Scan(&test.pname, &test.fname, &test.lname, &test.age)
	if err != nil {
		return nil, err
	}
	return &test, nil
}

func AddTest(test Test) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	query := "insert into test (pname,fname,lname,age) values (? , ? , ? ,?) "
	result, err := db.Exec(query, test.pname, test.fname, test.lname, test.age)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected() //เป็นการเช็คผลกระทบที่เกิดขึ้นกับข้อมูล
	if err != nil {
		return err
	}
	// จำนวนข้อมูลที่ได้รับผลกระทบด้วยคำสั่ง
	if affected <= 0 {
		panic("Cannot Insert")
	}
	return nil
}

func UpdateTest(test Test) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	query := "update test set pname=?,fname=?,lname=? where age=? "
	result, err := db.Exec(query, test.pname, test.fname, test.lname, test.age)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected() //เป็นการเช็คผลกระทบที่เกิดขึ้นกับข้อมูล
	if err != nil {
		return err
	}
	// จำนวนข้อมูลที่ได้รับผลกระทบด้วยคำสั่ง
	if affected <= 0 {
		panic("Cannot Update")
	}
	return nil
}

func DeleteTest(age int) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	query := "delete from test where age=? "
	result, err := db.Exec(query, age)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected() //เป็นการเช็คผลกระทบที่เกิดขึ้นกับข้อมูล
	if err != nil {
		return err
	}
	// จำนวนข้อมูลที่ได้รับผลกระทบด้วยคำสั่ง
	if affected <= 0 {
		panic("Cannot Delete")
	}
	return nil
}
