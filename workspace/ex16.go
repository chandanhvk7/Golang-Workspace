//pointers

package main

import "fmt"

func main(){
	n1:=1234
	fmt.Printf("Type of n1 is %T, n1=%v, address of n1= %v\n",n1,n1,&n1)

	p1:=&n1
	fmt.Printf("Type of p1 is %T, p1=%v, value referred by p1= %v\n",p1,p1,*p1)
	*p1=1000
	fmt.Printf("Type of n1 is %T, n1=%v, address of n1= %v\n",n1,n1,&n1)
}
