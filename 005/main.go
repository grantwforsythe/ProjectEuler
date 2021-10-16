package main

import "fmt"


func main() {
	num := 1
	nums := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for _, i := range nums {
		num = lcm(num, i)
	} 
	fmt.Printf("%d\n", num)
}


func lcm(a int, b int) int {
    /*
        Description:    Find the least common multiple between two numbers. 
        Parameters:     a               - int
                        b               - int
        Return:         the least common multiple 
        Example of Usage:
        >> lcm(3, 5)
        ans =
            15 
    */ 
	return a*b / gcd(a, b)
}


func gcd(a int, b int) int {
	/*
        Description:    Find the greatest common divisor between two numbers. 
        Parameters:     a               - int
                        b               - int
        Return:         the greatest common divisor
        Example of Usage:
        >> gcd(18, 24)
        ans =
            6
	*/
	if b == 0{
		return a
	} else {
		return gcd(b, a % b)
	}
}