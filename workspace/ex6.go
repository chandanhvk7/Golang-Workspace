package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args)==1 {
		fmt.Println("No args!!!")
		return
	}
	fmt.Println("Here are your arguments:")
	for index, value := range os.Args[1:] {
		fmt.Printf("%d->%v\n", index, value)
	}
	fmt.Println("Args without indices:")
	for _, value := range os.Args[1:] {
		fmt.Printf("%v\n", value)
	}
}
