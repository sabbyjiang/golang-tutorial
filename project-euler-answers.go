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

func main() {
	// Find the sum of all the multiples of 3 or 5 below 1000.
  fmt.Println(sumThreeOrFive(1000))
}