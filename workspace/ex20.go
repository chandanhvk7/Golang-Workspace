
package main

import "fmt"

type Product struct{
	ID int
	Name string
	Price float64
}


//Any changes made here does not affect the actual values
// If changes have to be made, change it to func(p *product) Print(){...}
func(p Product) Print(){
	p.Price+=500
	fmt.Printf("ID:%v\n",p.ID)
	fmt.Printf("Name:%v\n",p.Name)
	fmt.Printf("Price:%v\n",p.Price)
	fmt.Println()
}

func main(){
	p1:=Product{Name:"Logitech Mouse",Price:4500,ID:23}
	p2:=Product{Name:"Zebronyx Mouse",Price:500,ID:2}
	var p3 Product
	p1.Print()
	p2.Print()
	p3.Print()
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
