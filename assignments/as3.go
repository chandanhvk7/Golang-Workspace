package main

import ("fmt" 
"math")

func main() {
 var f,t int
 fmt.Print("Enter the range n,m:")
 fmt.Scan(&f,&t)
 sum:=sum0fPrimes(f,t)
 fmt.Printf("The sum of Primes between %d and %d is %d",f,t,sum)
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <=int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func sum0fPrimes(from, to int) int {
	var sum int
	sum = 0
	for i := from; i <= to; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}
