package main

import "fmt"

func fact(num int) int{
	var sum int =1;
	for num>0{
		sum*=num
		num--
	}
	return sum
}

func pow(num float64,x int) float64{
	var sum float64=1
	for x>0{
		sum*=num
		x--
	}
	return sum
}

func sine(x float64) float64{
	var sum float64=0
	var sin float64=-1
  for i:=1;i<10;i++{
	j:=(2*i)-1
	sin*=-1
	sum+=sin*(pow(x,j)/float64(fact(j)))
  }
  return sum
}

func main(){
	var deg float64;
	fmt.Print("Enter the angle in degrees:")
	fmt.Scan(&deg)
	rad:=3.14*deg/180
	fmt.Println("The sine of the angle is",sine(rad))
}