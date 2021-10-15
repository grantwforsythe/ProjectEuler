package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)


var prime_nums []int


func main() {
    // read the file
    file, err := os.Open("primes-to-100k.txt")

    if err != nil {
        log.Fatal(err)
    }

    // close the file
    defer file.Close()

    // convert each number into an int and append to a slice
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        num, _ := strconv.Atoi(scanner.Text())
        prime_nums = append(prime_nums, num)
    }

    if err := scanner.Err(); err != nil {
        log.Print(err)
    }

    var num int
    fmt.Print("Find the largest prime factors of a nonzero number: ")
    fmt.Scanf("%d", &num)

    result, err := prime_factorization(num)

    if err != nil {
        log.Print(err)
    } else {
        fmt.Printf("%d\n", result)
    }
}

func prime_factorization(num int) (int, error) {
    /*
        Description:    Find the largest prime factors of a nonzero number.
        Parameters:     num             - int
        Return:         that largest prime factor
        Example of Usage:
        >> prime_facorization(13195)
        ans =
            29
    */

    if num == 0 {
        return 0, errors.New("Math: num cannot be equal to 0")
    }

    for _, prime_num := range prime_nums {
        if num % prime_num == 0 {
            num /= prime_num
            if num == 1 {
                return prime_num, nil
            } else {
                prime_factorization(num)
            }
        }
    }
    return 0, nil
}
