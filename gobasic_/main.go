package main

import "github.com/nikepitchaya/hello/Sharepackage"
import "fmt"

func main() {
	e := Sharepackage.Players{
		pname: "nay",
		fname: "nike",
		lname: "good",
		id:    5,
	}
	e.LeaveandShow()
	fmt.Println(e)
}
