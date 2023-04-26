package main

import "golang.org/x/text/width"

type geometry interface {
	area()
	perim()
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() {
	return r.width * r.height
}
func (r rect) perim() {
	return 2 * r.width * 2 * r.height
}
func (c circle) area() {
	return math.Pi * c.radius * c.radius
}
func measure (g geometry){
	
}
func main() {
	rec := rect{width: 5, height: 10}
	cir := circle{radius: 15}
}
