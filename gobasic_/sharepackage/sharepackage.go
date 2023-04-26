package Sharepackage

import "fmt"

type Players struct {
	pname string
	fname string
	lname string
	id    int
}

func (p Players) LeaveandShow() {
	fmt.Printf("%s %s %s %d", p.pname, p.fname, p.lname, p.id)
}
