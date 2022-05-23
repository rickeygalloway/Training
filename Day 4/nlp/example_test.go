package nlp_test

/*
from command line
go test


*/

import (
	"fmt"
	"git.shipt.com/nlp"
)

func ExampleTokenize() {
	text := "Who's on first?"
	tokens := nlp.Tokenize(text)
	fmt.Println(tokens)

	//below return value is for an array

	//Output:
	// [who s on first]
}

func ExampleTestFunc() {
	text := nlp.TestFunc("testing")
	fmt.Println(text)

	//Output:
	// testing
}
