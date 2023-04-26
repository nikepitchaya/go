package main

import "fmt"

type products struct {
	id    int
	name  string
	price int
}

func main() {
	var rec products
	rec.id = 1
	rec.name = "Iphone"
	rec.price = 500
	fmt.Println(rec)
	show(rec)
	
	update(&rec)
	show(rec)

	rec = rec.clear()
	show(rec)

}
func show(r products) {
	fmt.Println(r)
}
func update(r *products) {
	r.id = r.id + 500
	r.name = "Keroro"
	r.price = r.price + 5000
}
// add function in struct
func (p products) clear() products{
	p.id=0
	p.name=""
	p.price=0
	return p
}
