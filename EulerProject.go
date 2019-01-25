package main

import (
	"fmt"
	// "math"
)




// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.
func Multiples (numbersBelow int, multiplesOf int) int {
	sum := 0
	for i := 0; i < numbersBelow; i++ {
		if i%multiplesOf  == 0 {
			sum += i
		}
	}
	return sum
}

// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

func Fibonacci (upperBorder int) ([]int, int) {
	arr := []int{1, 2}
	el := 0
	sum := 2
	for i := 1; el <= upperBorder; i++ {
		el = ( arr[i] + arr[i-1] )
		if el < upperBorder {
			arr = append(arr, el)
			if el%2 == 0 {
				sum += el
			}
		}
	}
	return arr, sum
}

// The prime factors of 13195 are 5, 7, 13 and 29
// What is the largest prime factor of the number 600851475143

func PrimeFactors(n int64) (pfs []int64) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}
	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; int64(i*i) <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%int64(i) == 0 {
			pfs = append(pfs, int64(i) )
			n = n / int64(i)
		}
	}
	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func main () {
	// sum := (Multiples(1000, 3) + Multiples(1000, 5) - Multiples(1000, 3*5) )
	// fmt.Print( Fibonacci(1000000*4) )
	fmt.Println(PrimeFactors(600851475143) )
}