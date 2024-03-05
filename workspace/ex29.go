package main

import (
	"fmt"
	"time"
)

func go1(ch chan string){
	for{
		ch<-"channel1"
		time.Sleep(500*time.Millisecond)
	}
}
func go2(ch chan string){
	for{
		ch<-"channel2"
		time.Sleep(250*time.Millisecond)
	}
}
func main(){
	ch1:=make(chan string,10)
	ch2:=make(chan string,10)
	fmt.Println("Starting go1...")
	go go1(ch1)
	fmt.Println("Starting go2...")
	go go2(ch2)
	for{
		select{
		case channel1Data:=<-ch1:
			fmt.Println("Got data from",channel1Data)
		case channel2Data:=<-ch2:
			fmt.Println("Got data from",channel2Data)
		}
	}
}