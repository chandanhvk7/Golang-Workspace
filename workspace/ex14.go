//Arrays 
package main

import "fmt"

func main(){
	num:= [5]int{10,20,30,40,50}
	fmt.Println("Nums is",num)
	fmt.Print("Enter nums values:")
	for i:=0;i<len(num);i++{
		fmt.Scan(&num[i])
	}
	fmt.Println("Nums is",num)
	for index,value:=range(num){
		fmt.Printf("nums[%d] is %d\n",index,value)
	}
}