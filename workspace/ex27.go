//Buffered channel usage
package main

import "fmt"

var ch chan Factorial

type Factorial struct {
	n int
	f int
}

func findFact(num int){
	f:=1
	for i:=1;i<=num;i++{
		f*=i
	}
	ch<-Factorial{num,f}
}

func main(){
	ch=make(chan Factorial,1)
	nums:=[]int{10,2,4,3,2,5,7}
	for _,n:=range nums{
		go findFact(n)
	}
	for i:=0;i<len(nums);i++{
		obj:=<-ch
		fmt.Printf("Factorial of %d is %d\n",obj.n,obj.f)
	}
}