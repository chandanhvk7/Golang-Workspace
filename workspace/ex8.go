package main

import "fmt"

func main(){
	a,b:=subtotal(153,36,39,25,36)
	fmt.Printf("a=%v,b=%v\n",a,b)
	fmt.Printf("a=%T,b=%T\n",a,b)
}

func subtotal(nums ...int)(sum int, avg float64){
	for _,n:=range nums{
		sum+=n
	}
	avg=float64 (sum)/float64(len(nums))
	return
}