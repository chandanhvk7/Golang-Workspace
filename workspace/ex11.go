//Deferred Statement
package main

import "fmt"

func main(){
	n:=10
	fmt.Println("Using Deferred statements...")
	fmt.Printf("n=%v\n",n)
	n++
	defer fmt.Printf("Statement 1,n=%v\n",n)
	n++	
	defer fmt.Printf("Statement 2,n=%v\n",n)	
	n++
	defer fmt.Printf("Statement 3,n=%v\n",n)	
	fmt.Printf("End of main(),n=%v\n",n)
}