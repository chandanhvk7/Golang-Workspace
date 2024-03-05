package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("1.Add new customer\n2:View all customers\n3:Search customers by city\n4:Delete a customer (by id)\n5:Search customer by id and edit/update details\n6:Exit")
	var choice int
	fmt.Scan(&choice)
	for{
		switch choice {
		case 1:
			addNewCus()
		case 2:
			viewAllCus()
		case 3:
			searchCus()
		case 4:
			delCus()
		case 5:
			editCus()
		case 6:
			break
		}
	}
}