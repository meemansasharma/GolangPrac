package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	hieght float64
}

type square struct {
	side float64
}

func (tri triangle) getArea() float64 {

	return 0.5 * tri.base * tri.hieght
}

func (sq square) getArea() float64 {

	return sq.side * sq.side
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func getshapes() {
	tri := triangle{
		base:   10,
		hieght: 15,
	}
	sq := square{
		side: 9,
	}
	printArea(tri)
	printArea(sq)
}
