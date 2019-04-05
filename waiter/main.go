package main

import "fmt"
import "time"

func main(){
	fmt.Println("Waiter started")	
	defer fmt.Println("Waiter finished")
	select {
		case <-time.After(10 * time.Second):
	}	

}