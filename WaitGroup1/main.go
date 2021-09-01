package main

import (
	"fmt"
	"time"
)

func main() {

	var data int

	go func() {

		data++
	}()

	if data == 0 {

		fmt.Printf("The value is %v\n", data)
	}

	time.Sleep(10 * time.Millisecond)
}
