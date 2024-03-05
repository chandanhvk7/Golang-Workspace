//Functions assigned to variables
package main

import "fmt"

func main(){
	greet:=func(){
		fmt.Println("Hello World!")
	}
	fmt.Printf("The type of greet is %T\n",greet)
	greet()
	welcome:=greet
	welcome()
	func(){
		fmt.Println("This is an example of an anonymous func")
	}()
}