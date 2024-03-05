//Convert String to Int
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0.
	var nonum string
	for _, arg := range os.Args[1:] {
		if n, err := strconv.ParseFloat(arg, 64); err == nil {
			sum += n
		} else {
			nonum+=arg+","
		}
	}
	fmt.Printf("Sum of all the input in command line is %v\n",sum)
	fmt.Println("You entered non numerical inputs-",nonum)
}
