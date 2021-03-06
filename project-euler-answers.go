package main

import (
  "fmt"
)

func threeOrFive(x int) (yes bool) {
  check := x % 3 == 0 || x % 5 == 0
  return check
}

func sumThreeOrFive(limit int) (int) {
	sum := 0
	for i:= 0 ; i < limit ; i++ {
		if threeOrFive(i){
		  sum += i;
		}
	}
	return sum
}

func fib(n int) (int) {
	a:= 0
	b:= 1
	c:= 0

	for i := 0 ; i < n ; i++ {
		c = a + b
		a = b
		b = c
	}
	
	return c
}

func sumOfEvenFib(n int) (int) {
	sum := 0
	
	for i := 0 ; fib(i) < n ; i++ {
		if fib(i) % 2 == 0 {
			sum += fib(i);
		}
	}
	return sum
}

func main() {
	// Find the sum of all the multiples of 3 or 5 below 1000.
  fmt.Println(sumThreeOrFive(1000))

	/** Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
	1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
	By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms. */

	fmt.Println(sumOfEvenFib(4000000))
}