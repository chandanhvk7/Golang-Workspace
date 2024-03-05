//garbage collector

package main

import "fmt"

func getNames() *[]string{
	name:=[]string{"Chandan","Kshatriya","Hosanagara","Venkatesh"}
	return &name
}
func someShit(){
	someNames:=getNames()
	fmt.Println("Names are",*someNames)
	fmt.Printf("The address of someNames is %v\n",&(*someNames)[0])
	//someNames is a local variable and will be removed at this time and the allocated heap memory is going
	//to be garbage collected since this is the only pointer variable to that loaction
}

func main(){
	someShit()
	someShit()
	someShit()
}
