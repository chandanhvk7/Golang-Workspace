package main

import "fmt"

func isEven(num int) bool {
	return num%2 == 0
}

func maxDays(month, year int) int {
	days := 31
	switch month {
	case 4, 6, 9, 11:
		days = 30
	case 2:
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			days = 29
		} else {
			days = 28
		}
	}
	return days
}

func main() {
	var mon,year int
	fmt.Print("Enter the Month and year:")
	fmt.Scan(&mon,&year)
	d:=maxDays(mon,year)
	fmt.Println(d)
}
