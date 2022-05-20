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

	pl1 := Player{
		Name:  "Rickey",
		Point: Point{10, 10},
	}
	fmt.Printf("pl1: %#v\n", pl1)
	fmt.Printf("pl1.X: %#v\n", pl1.Point.X)
	pl1.Move(0, 0)
	fmt.Printf("pl1.X (moved): %#v\n", pl1.Point.X)

	ms := []Mover{
		&p1,
		&pl1,
		p3,
	}
	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Printf("%#v\n", m)
	}

	k := Copper
	fmt.Println("k:", k)
}

//implement fmt.Stringer inferface
func (k Key) String() string {
	switch k {
	case Copper:
		return "Copper"
	case Crystal:
		return "Chrystal"
	case Jade:
		return "Jade"
	}
	// Warning: Don't use %s or %v here (∞ recursion)
	return fmt.Sprintf("<key %d>", k)
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

type Player struct {
	Name string
	Keys []Key
	//X float64 //This will shadow Point.X
	Point //Player is embedding Point
}
type Point struct {
	X int
	Y int
}

func moveAll(ms []Mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

type Mover interface {
	Move(int, int)
}

// iota is automatically incremented by 1 in every const group
const (
	Crystal Key = iota + 1 // 1 << iota (bitmask)
	Copper
	Jade
)

type Key uint8

// Rule of thumb: Return types, accept interfaces
//functions vs methods

//Whats the minimum interface required for sorting something
/*

type Sortable interface{
	Len() int				// need to know the length
	Less(i,j int) bool		//
	Swap(i,j int) bool
}
The interface here matches
https://pkg.go.dev/sort#Interface


*/
