/*
Given two arrays of integers, determine all integers that satisfy the following
two conditions:
1. The elements of the first array are all factors of the integer being considered
2. The integer being considered is a factor of all elements of the second array

These numbers are referred to as being between the two arrays. Determine how many
such numbers exist.

Example
a = [2, 6]
b = [24, 36]

There are two numbers between the arrays: 6 and 12.
For the first value:
	6 % 2 = 0
	6 % 6 = 0
	24 % 6 = 0
	36 % 6 = 0
For the second value:
	12 % 2 = 0
	12 % 6 = 0
	24 % 12 = 0
	36 % 12 = 0
=> Returns 2

Function Parameters:
- int a[n]: an array of integers
- int b[m]: an array of integers

Returns:
- int: the number of integers that are between the sets

Input Format:
- The first line contains two space-separated integers, n and m, the number
of elements in arrays a and b.
- The second line contains n distinct space-separated integers a[i] where 0 < i < n.
- The third line contains m distinct space-separated integers b[j] where 0 < j < m.

Sample Input:
$ 2 3
$ 2 4
$ 16 32 96
Sample Output:
$ 3
Explanation:
2 and 4 divide evenly into 4, 8, 12 and 16.
4, 8 and 16 divide evenly into 16, 32, 96.
4, 8 and 16 are the only three numbers for which each element of a is a factor and each
is a factor of all elements of b.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Function to find the Greatest Common Divisor (GCD)
func gcd(a, b int32) int32 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to find Least Common Multiple (LCM)
func lcm(a, b int32) int32 {
	return a * b / gcd(a, b)
}

func getTotalX(a []int32, b []int32) int32 {
	lcmValue := a[0]
	for _, value := range a[1:] {
		lcmValue = lcm(lcmValue, value)
	}

	gcdValue := b[0]
	for _, value := range b[1:] {
		gcdValue = gcd(gcdValue, value)
	}

	var count int32 = 0
	for i := lcmValue; i <= gcdValue; i += lcmValue {
		if gcdValue%i == 0 {
			count++
		}
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var brr []int32

	for i := 0; i < int(m); i++ {
		brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
		checkError(err)
		brrItem := int32(brrItemTemp)
		brr = append(brr, brrItem)
	}

	fmt.Println(getTotalX(arr, brr))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
