package main

import "fmt"




func genMsg(ch1 chan<- string){

	//send message on ch1

	ch1 <- "welcome"

}


func relayMsg(ch1 <-chan string,ch2 chan<- string){

	//recv message on ch1
	//send it on ch2
  
	m := <-ch1
	ch2 <- m
}

func main(){


	ch1:= make(chan string)
	ch2:= make(chan string)


	go genMsg(ch1)

	go relayMsg(ch1,ch2)


	v := <-ch2

	fmt.Printf("%v",v)

	   
}