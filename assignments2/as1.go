package main

import (
	"fmt"
	"errors"
)

func numberToWords(num int) (string, error) {
    if num<0{
		return "", errors.New("This func cannot handle negative things like you")
	}
	var str string = ""
	str += convert(num/10000000, " crore ")
	str += convert((num/100000)%100, " lakh ")
	str += convert((num/1000)%100, " thousand ")
	str += convert((num/100)%10, " hundred ")
	str += convert(num%100, "")
	if num == 0 {
		str = "zero"
	}
	return str,nil
}

func convert(num int, s string) string {
	ten := []string{"", "", "twenty ", "thirty ", "forty ", "fifty ", "sixty ", "seventy ", "eighty ", "ninety "}
	one := []string{"", "one ", "two ", "three ", "four ", "five ", "six ", "seven ", "eight ", "nine ", "ten ", "eleven ", "twelve ", "thirteen ", "fourteen ", "fifteen ", "sixteen ", "seventeen ", "eighteen ", "nineteen "}
	var str string = ""
	if num > 19 {
		str += ten[num/10] + one[num%10]
	} else {
		str += one[num]
	}
	if num != 0 {
		str += s
	}
	return str
}

func main() {
	var n int
	fmt.Print("Enter the number:")
	fmt.Scan(&n)
	if str,err:=numberToWords(n);err==nil{
		fmt.Println(str)
	}else{
		fmt.Println("There was an error!",err.Error())
	}
}
