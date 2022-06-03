package main

/*
  #include "hello.h"
*/
import "C"
import (
	"fmt"
)

func main() {

	fmt.Println("TESTING")

	//Call to void function without params
	err := C.Hello()
	if err != nil {
		fmt.Println(err)
	}
}
