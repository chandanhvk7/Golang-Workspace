package main

import "fmt"

func dayNumber (day, month, year int) int{
 
    var t [12]int = [12]int{ 0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4 };
    if month < 3{
		year--
	}
    return ( year + year/4 - year/100 + year/400 + t[month-1] + day) % 7;
}
func main(){
	var month,year int
	fmt.Print("Enter the month and year:")
	fmt.Scan(&month,&year)
	printCalendar(month,year)
}
func calDays(month,year int) int{
	var days int
	switch month{
	case 1,3,5,7,8,10,12:
		days=31
	case 4,6,9,11:
		days=30
	case 2:
		if year<0{
			return 0
		}
		if isLeap(year){
			days=29
		}else{
			days=28
		}
	}
	return days
}
func isLeap(year int) bool{
	return year%400==0 || (year%4==0 && year%100!=0)
}
func printCalendar(month, year int) {
	firstDay:=dayNumber(1,month,year)
	fmt.Println("Su Mo Tu We Th Fr Sa")
	days:=calDays(month,year)
	for i:=0;i<firstDay;i++{
		fmt.Print("   ")
	}
	for i:=1;i<=days;i++{
		if (i-1+firstDay)%7==0{
			fmt.Println()
		}
		fmt.Printf("%2d ",i)
	}
}
