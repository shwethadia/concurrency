package main

import "fmt"

//How to get value computed by goroutine into main routine

func main(){

		//	var ch chan int

		ch := make(chan int)

			go func(a,b int){

					c := a+b
					ch <- c
			}(1,2)

			r := <-ch
			fmt.Printf("Computed value %v\n",r)
		
}