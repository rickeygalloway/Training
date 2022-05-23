package main

import (
	"fmt"
	"sync"
)

func main() {
	p := Payment{
		From:   "WEC",
		To:     "ACME",
		Amount: 123.45,
	}

	p.Execute()
	p.Execute()
}
func (p *Payment) Execute() {
	p.once.Do(p.execute)
}

func (p *Payment) execute() {
	fmt.Printf("%s -> %.f$-> %s\n", p.From, p.Amount, p.To)
}

type Payment struct {
	From   string
	To     string
	Amount float64

	once sync.Once
}
