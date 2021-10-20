package main

import (
    "strconv"
    "fmt"
)


func main() {
    result := 0 

    for a := 100; a < 1000; a++ {
        for b := 100; b < 1000; b++ {
            num := a * b
            str_num := strconv.Itoa(num)
            if ispalindrome(str_num) && (num > result) {
                result = num
            }
        }
    }
	fmt.Printf("Answer: %d\n", result)
}


func ispalindrome(num string) bool {
    /*
        Description:    Check if a string is a panlindrome.
        Parameters:     num             - string
        Return:         a boolean value 
        Example of Usage:
        >> ispalindrome("racecar")
        ans =
            True 
        >> ispalindrome("dog")
        ans =
            False
    */
    return num == reverse(num)
}


func reverse(str string) string {
    /*
        Description:    Reverses a string.
        Parameters:     str             - string
        Return:         reversed string 
        Example of Usage:
        >> reverse("cat")
        ans =
            tac 
    */
	r := []rune(str)
    var res []rune
    for i := len(r) - 1; i >= 0; i-- {
        res = append(res, r[i])
    }
	return string(res)
}
