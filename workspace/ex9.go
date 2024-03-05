//ERROR HANDLING

package main

import (
	"fmt"
	"errors"
)

func factorial(num int)(int ,error) {
	if num<0{
		return 0, errors.New("This func cannot handle negative things like you")
	}
	f:=1
	for num>1{
		f*=num
		num--
	}
	return f,nil
}

func main(){
	fmt.Println("Error Handling......")
	var n int
	fmt.Print("Enter a number:")
	fmt.Scan(&n)
	if fact,err:=factorial(n);err!=nil{
		fmt.Println("There was an error:",err.Error())
	}else{
		fmt.Printf("Factorial of %v is %v\n",n,fact)
	}
}