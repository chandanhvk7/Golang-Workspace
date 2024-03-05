//pointers 2

package main

import "fmt"

func increment(n* int){
	*n+=10
	fmt.Printf("The value and addr of n in increment() is %v and %v\n",*n,n)
}
func main(){
	num:=100
	fmt.Printf("num=%v,addr of num is %v\n",num,&num)
	increment(&num)
	fmt.Printf("num=%v,addr of num is %v\n",num,&num)
	increment(&num)
	fmt.Printf("num=%v,addr of num is %v\n",num,&num)

}
