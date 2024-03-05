//Interfaces and polymorphic func
package main

import "fmt"

//Interface
type Shape interface{
	Area() float64
	GetName() string
}

type Circle struct{
	Radius float64
}

func(c Circle) Area() float64{
	return c.Radius*c.Radius*3.14
}

func(c Circle) GetName() string{
	return "Circle"
}

type Traingle struct{
	Height float64
	Base float64
}

func(t Traingle) Area() float64{
	return t.Base*0.5*t.Height
}

func(t Traingle) GetName() string{
	return "Triangle"
}

//polymorfic func
func ProcessShape(c1 Shape){
	fmt.Printf("Area of the %v is %v\n",c1.GetName(),c1.Area())
}

func main(){
	c1:=Circle{5}
	t1:=Traingle{2,3}
	ProcessShape(c1)
	ProcessShape(t1)
}
