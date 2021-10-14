package main

import "fmt"


const n int = 4000000


func fib() int {
    /*
        Description:    Compute the Fibonacci Sequence.
        Parameters: 	none
        Return:	        the sum of the even-valued terms in the Fibonacci Sequence whose values
                        do not exceed 4 million
    */

    result := 0
    // starting values
    a := 0
    b := 1
    for a+b < n {
        temp := a + b
        if temp%2 == 0 {
            result += temp
        }
        a = b
        b = temp
    }
    return result
}


func main() {
    fmt.Println("Answer: ", fib())
}
