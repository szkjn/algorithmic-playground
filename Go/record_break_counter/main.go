/*
Write a program to analyze a list of game scores and determine how many times
a player breaks their highest and lowest score records. The first input is an
integer representing the number of scores. The second input is a space-separated
list of integers representing the scores.

The program should output two integers: the first representing the count of new
high score records, and the second representing the count of new low score records.
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

func recordBreakCounter(scores []int32) []int32 {
	res := []int32{0, 0}
	var most, least int32 = 0, 0

	for idx, score := range scores {
		if idx == 0 {
			most = score
			least = score
		} else {
			if score > most {
				most = score
				res[0]++
			}
			if score < least {
				least = score
				res[1]++
			}
		}
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	scoresTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var scores []int32

	for i := 0; i < int(n); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	result := recordBreakCounter(scores)

	fmt.Printf("%d %d", result[0], result[1])
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
