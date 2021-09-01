package main

import (
	"fmt"
	"time"
)

func fun(s string) {

	for i := 0; i < 3; i++ {

		fmt.Println(s)
		time.Sleep(10 * time.Millisecond)
	
	}
}

func main() {

	fun("Direct Call")


	//TODO: write  goroutine with different variants for function call



	//goroutine function call

	go fun("goroutine")

	//goroutine with anonymous function

	go func(){

		fun("Anonymous Goroutine")
	}();

	//goroutine with function value call

	fv := fun
	go fv("function value goroutine")

	//wait for the goroutine to end

	time.Sleep(100 * time.Millisecond)

	fmt.Print("\nDONE\n")

}