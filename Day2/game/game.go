package main

import (
	"fmt"
	"log"
)

func main() {
	var p0 Point
	fmt.Println(p0)
	fmt.Printf("%#v\n", p0)

	p1 := Point{1, 1}
	fmt.Printf("%#v\n", p1)

	p2 := Point{Y: 2}
	fmt.Printf("%#v\n", p2)

	p3, err := NewPoint(10, 20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("p3: %#v\n", p3)

	p4 := Point{-1, 40}
	p4.Move(10, 20)
	fmt.Printf("p4 (move): %#v\n", p4)

	p3.Move(0, 0)
	fmt.Printf("p3 (move): %#v\n", p3)

}

// most of the time, use pointer receiver
func (p *Point) Move(x, y int) {
	p.X = x
	p.Y = y
}

func NewPoint(x, y int) (*Point, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds (%d/%d)", x, y, maxX, maxY)
	}
	p := Point{x, y}
	return &p, nil
}

const (
	maxX = 100
	maxY = 600
)

type Point struct {
	X int
	Y int
}
