package main

import (
	"fmt"
	"math"
)


func main(){
	limit := 100
	sum := limit * (limit + 1) /2
	sum_sq := ((2 * limit + 1)*(limit + 1) * limit) / 6
	
	result := int(math.Pow(float64(sum), 2) - math.Pow(float64(sum_sq), 2))
	fmt.Printf("Answer: %f", result)
	
}