// Analysing Buffered Channels
package main

import (
	"fmt"
	"time"
)

func main(){
	ch:=make(chan string,3)
	go func(){
		time.Sleep(5*time.Second)
		fmt.Println(<-ch)
		fmt.Println("5 second Go is done")
	}()
	go func(){
		time.Sleep(10*time.Second)
		fmt.Println("10 second Go is done")
		fmt.Println(<-ch)
	}()
	ch<-"chvk1"
	ch<-"chvk2"
	ch<-"chvk3"
	ch<-"chvk4"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch<-"chvk5"
	ch<-"chvk6"

	fmt.Println(<-ch)
	fmt.Scanln()
}