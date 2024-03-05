//Deferred function
package main

import "fmt"

func doSomething(){
	defer func(){
		if r:=recover();r!=nil{
			fmt.Println(r)
		}
	}()
	var n,d int
	fmt.Print("Enter Numerator:")
	fmt.Scan(&n)

	fmt.Print("Emter Denominator:")
	fmt.Scan(&d)

	q:=n/d
	fmt.Printf("%v/%v is %v\n",n,d,q)
}

func main(){
	fmt.Println("Start of main...")
	doSomething()
	fmt.Print("End of main...")
}