//GETTING ENV VARIABLES
package main

import (
	"fmt"
	"os"
)

func main() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	fmt.Println("--------------------------------------------------------")
	path:=os.Getenv("PATH")//can get any variable this way (VINOD SIR:9731424784)
	opersys:=os.Getenv("OS")
	fmt.Println("Path is",path)
	fmt.Println("OS:",opersys)
}
