package main

import (
	"fmt"
	"sync"
)





func main(){


		var data int
		var wg sync.WaitGroup
		wg.Add(1)

		go func(){

				defer wg.Done()
				data++
		}()


		wg.Wait()
		if data == 0 {

				fmt.Print("The value of data is 0\n")
		}else{

			fmt.Printf("The value of data is %v\n",data)

		}
		fmt.Print("\nDone\n")
}