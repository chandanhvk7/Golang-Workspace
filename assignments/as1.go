package main

import "fmt"
func isValidDate(year,month,day int) bool{
	if year<0{
		return false
	}
    if month <0 || month>12 {
		return false
	}
	if day<0{
		return false
	}
	var days int=31
	switch month{
	case 4,6,9,11:
		days=30
	case 2:
		if year%400==0 || (year%4==0 && year%100!=0){
			days=29
		}else{
			days=28
		}
	}
	if day>days {
		return false
	}
	return true
}

func main(){
	var d,m,y int
	fmt.Print("Enter the date in dd/mm/yyyy format:")
	fmt.Scanf("%d/%d/%d",&d,&m,&y)
	tf:=isValidDate(y,m,d)
	fmt.Print(tf)
}