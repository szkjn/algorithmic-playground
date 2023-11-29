/*
Write a Go function, migratoryBirds, that takes an array of bird sightings
where each element is a bird type ID. Your task is to determine the ID of
the most frequently sighted bird type. If more than one type has the
maximum sightings, return the ID of the smallest type. The bird types are
guaranteed to be integers 1 through 5.
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

func migratoryBirds(arr []int32) int32 {
	frequency := make(map[int32]int32)
	var maxCount, maxType int32 = 0, int32(6)

	for _, bird := range arr {
		frequency[bird]++
		if frequency[bird] > maxCount || (frequency[bird] == maxCount && bird < maxType) {
			maxCount = frequency[bird]
			maxType = bird
		}
	}

	return maxType
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := migratoryBirds(arr)

	fmt.Println(result)
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
