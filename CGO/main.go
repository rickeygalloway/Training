package main

// #include "mathlib.h"
import "C"
import "fmt"

func main() {

	added := C.Add(C.uint32_t(10), C.uint32_t(10))
	//added := 20
	fmt.Printf("10+10 = %d!\n", added)
}
