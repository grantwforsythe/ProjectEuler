package main

import "fmt"

// max number
const n int = 1000

func main() {
	fmt.Println("Answer: ", multiples_3or5())
}

func multiples_3or5() int {
	/*
	   Description:    Finds the sum of all the multiples of 3 and 5 less than n.
	   Parameters:     none
	   Return:         the sum of all natural numbers less than n that are divisible by 3 or 5
	   Example of Usage:
	   >> const n int = 10
	   >> result := multiples_3or5()
	   ans =
	       23
	*/
	result := 0
	for num := 1; num < n; num++ {
		if num%3 == 0 || num%5 == 0 {
			result += num
		}
	}
	return result
}
