package main

import (
	"fmt"
	"math"
)


func main(){
	limit := 100
	sum := limit * (limit + 1) /2
	// proof: https://github.com/grantwforsythe/project-euler/tree/main/proofs/006.pdf
	sum_sq := ((2 * limit + 1)*(limit + 1) * limit) / 6
	
	result := math.Pow(float64(sum), 2) - float64(sum_sq)
	fmt.Printf("Answer: %d\n", int(result))
	
}