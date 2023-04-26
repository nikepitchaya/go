package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	var company employee
	company.employeeID="NIKE31"
	company.employeeName="ROUGE"
	company.phone="091"
	fmt.Println(company)

	employee1 := employee{
		employeeID:   "101",
		employeeName: "202",
		phone:        "303",
	}
	fmt.Println(employee1)

	office1 := [3]string{}
	office1[0]=employee1.employeeID
	fmt.Println(office1)

	office2 := [3]string{}
	office2[0] = "wod"
	fmt.Println(office2)

	office3 := [3]employee{}
	office3[0]= employee{
		employeeID:   "101",
		employeeName: "201",
		phone:        "301",
	}
	office3[1]= employee{
		employeeID:   "103",
		employeeName: "203",
		phone:        "303",
	}
	office3[2]= employee{
		employeeID:   "104",
		employeeName: "204",
		phone:        "304",
	}
	fmt.Println(office3)

	officeList := []employee{}
	officeList = append(officeList,office3[0] )
	officeList = append(officeList,office3[1] )
	officeList = append(officeList,office3[2] )
	fmt.Println(officeList)
}