/*
Interview question of the week by Cassidy Williams (issue #273) :

    Given a positive integer n, return all of its anti-divisors. 
    Anti-divisors are numbers that do not divide a number by the 
    largest possible margin (1 is not an anti-divisor of any number). 
    More information here : https://oeis.org/A066272/a066272a.html
*/

package main

import (
	"fmt"
)

func isAntiDivisor(k, n int) bool {
	if k%2 == 0 {
		return n%k == k/2
	}
	return n%k == (k-1)/2 || n%k == (k+1)/2
}

func antiDivisors(n int) []int {
	res := []int{}

	for k := 2; k < n; k++ {
		if isAntiDivisor(k, n) {
			res = append(res, k)
		}
	}

	return res
}

func main() {
	fmt.Println(antiDivisors(234))
}
