package main

import(
	"fmt"
	"os"
)

func main(){
	if len(os.Args)>1{
		fmt.Printf("Hello %s!",os.Args[1])
	}else{
		fmt.Printf("Hello User without a name!")
	}
}