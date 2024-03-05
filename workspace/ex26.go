//channels
package main

import (
	"fmt"
	"time"
)

func factorial(num int,ch chan int){
	f:=1
	for num>1{
		f*=num
		num--
	}
	time.Sleep(3*time.Second)
	ch<-f
}

func main(){
	ch:=make(chan int)
	fmt.Println("Creating a new Goroutine")
	go factorial(5,ch)
	fmt.Println("new GoRoutine created")
	fmt.Println("Waiting to receive")
	fact:=<-ch
	fmt.Println("Go a value from channel")
	fmt.Printf("factorial of 5 is %d\n",fact)
}