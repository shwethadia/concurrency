package main

import "fmt"



func main(){


	ch := make(chan int,6)
			go func(){

				for i:=0; i <6;i++{

					ch <- i

				}

				close(ch)
			}()


			for v := range ch {

				  fmt.Println(v)
			}
}