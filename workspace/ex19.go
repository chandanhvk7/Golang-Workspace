//struct

package main

import "fmt"

type Person struct{
	Name string
	Age int
	Contact Contact
}

type Contact struct{
	Email string
	Phone string
}
func main(){
	var p1 Person
	p2:=Person{Name:"Chandan",Age:21}
	p2.Contact.Email="chandan@gmail.com"
	p2.Contact.Phone="93753675122"
	fmt.Println("P1 is:",p1)
	fmt.Println("P2 is:",p2)
	ptr:=&p1
	(*ptr).Name="Batman"
	(*ptr).Age=972863
	(*ptr).Contact.Email="Vengeace@bat.com"
	(*ptr).Contact.Phone="69696969669"
	fmt.Println("P1 is",p1)
}
