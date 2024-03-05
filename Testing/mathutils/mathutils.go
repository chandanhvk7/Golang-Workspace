package mathutils

import "errors"

func Factorial(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("cannot calculate factorial of -ve number")
	}
	if num==0{
		return 1,nil
	}
	fact:=1
	for i:=1;i<=num;i++{
		fact*=i
	}
	return fact,nil
}