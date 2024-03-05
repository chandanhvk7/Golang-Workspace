package main

import ("fmt"
"math")

func main(){
	var n int
	fmt.Print("Enter a number:")
	fmt.Scan(&n)
	if isPrime(n){
		fmt.Println("The given number is Prime")
	}else{
		fmt.Println("The Given number is non prime")
	}
}

func isPrime(num int) bool {
	if num<2{
		return false
	}
	for i:=2;i<=int(math.Sqrt(float64(num)));i++{
		if num%i==0{
			return false
		}
	}
	return true
}