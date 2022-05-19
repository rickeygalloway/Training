package main

import (
	"fmt"
	"math"
	"os"
	"unicode/utf8"
)

func main() {
	s := "«ABC»"
	fmt.Println("s:", s)
	fmt.Println("len:", len(s))
	fmt.Println("len (runes):", utf8.RuneCountInString(s))

	c := s[3]
	fmt.Println("c:", c)
	fmt.Printf("c %c (%T)\n", c, c) // %c - rune, %T - type

	for i, r := range s {
		fmt.Printf("%c at %d\n", r, i)
	}

	fmt.Println("pi:", math.Pi)
	fmt.Printf("pi: %.2f\n", math.Pi)

	a, b := 1, "1"
	fmt.Printf("a=%v, b=%v\n", a, b)
	fmt.Printf("a=%#v, b=%#v\n", a, b)

	fmt.Fprintln(os.Stderr, "intro stderr")
	msg := fmt.Sprintf("a=%#v, b=%#v", a, b)
	fmt.Println(msg)
}
