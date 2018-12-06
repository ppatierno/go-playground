package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// note: interliving results between goroutines depends on how many core threads the goroutines are scheduled
	//fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))

	fmt.Scanln()
	fmt.Println("done")
}
