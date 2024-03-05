package main
import "fmt"

func  main(){
	name:="Chandan"
	fmt.Print(name)
	name="100"
	fmt.Print(" ",name)
	a,b:=100,200
	fmt.Print("\n")
	fmt.Print(a,b)
	var pi int=3
	var tf bool
	fmt.Print("\n")
	fmt.Printf("%T %d\n",pi,pi)
	fmt.Printf("%T",tf)
	fmt.Print(" ",tf)
}