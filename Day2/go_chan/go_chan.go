package main

import (
	"fmt"
	"time"
)

func main() {
	//go func() {
	//	fmt.Println("From a goroutine... Anonymous function")
	//}()

	go fmt.Println("From a goroutine")

	fmt.Println("From main ")

	for i := 0; i < 3; i++ {
		/* BUG : all goroutines reference i in line 17
		go func() {
			fmt.Println(i)
		}()
		*/
		//solution: pass arg
		i := i //this will shadow i from line 17
		go func(n int) {
			fmt.Println(n)
		}(i)
	}

	//The Go runtime does not wait for goroutines other than the main
	time.Sleep(10 * time.Millisecond)

	ch := make(chan string)
	go func() {
		ch <- "hi" //send		//This will create this issue: fatal error: all goroutines are asleep - deadlock! if not wrapped in goroutine
	}()

	val := <-ch //receive
	fmt.Println(val)

	values := []int{7, 2, 1, 8}
	fmt.Println(sleepSort(values))

}

//sleep sort
//for every n in values, spin a goroutine that will
//	sleep in n milliseconds
//	send n over a channel
// in the function body, collect values back from the channel to a slice and return it

func sleepSort(values []int) []int {
	ch := make(chan int)

	for _, v := range values {

		go func(n int) {
			time.Sleep(time.Duration(n) * time.Millisecond) //time.duration needed because type differences
			ch <- n
		}(v)
	}

	//not collect what was sent to the chanel
	var out []int
	for range values {
		n := <-ch
		out = append(out, n)
	}

	return out
	//for i := 0; i < len(values); i++ {
	//	go func(n int) {
	//		time.Sleep(n * time.Millisecond)
	//		ch <- values[i]
	//	}(i)
	//}
}

/*

Fun fact about me (and my family) is that in 2019, my wife, daughter, our maltipoo Caesar, and I moved to Cambodia for a year.

*/
