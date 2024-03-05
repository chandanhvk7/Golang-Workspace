//Slices

package main

import "fmt"

func main(){
	num:=[]int{10,20,30}
	fmt.Printf("len(nums)=%v\n",len(num))
	fmt.Printf("cap(nums)=%v\n",cap(num))

	fmt.Printf("Address of nums is %#x\n",&num[0])
	fmt.Println(num)
	num=append(num,0)
	fmt.Printf("Address of nums is %#x\n",&num[0])
	fmt.Println(num)
}