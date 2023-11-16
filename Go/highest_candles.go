/*
You are in charge of the cake for a child's birthday. The cake will have one candle for
each year of their total age. They will only be able to blow out the tallest of the
candles. Count how many candles are tallest.

Example:
- candles = [4, 4, 1, 3]

The maximum height candles are 4 units high. There are 2 of them, so return 2.

Function Parameter:
- int candles[n]: an array of candles heights

Returns:
- int: number of candles that are tallest

Input Format:
- The first line contains a single integer, n, the number (size) of candles.
- The second line contains n space-separated integers, where each integer i describes
the height of candles[i].
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func birthdayCakeCandles(candles []int32) int32 {
	var res, max_height int32 = 0, 0
	for _, el := range candles {
		if el > max_height {
			max_height = el
		}
	}
	for _, el := range candles {
		if el == max_height {
			res++
		}
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	fmt.Println("Enter total number of candles:")
	candlesCountTmp, _ := reader.ReadString('\n')
	candlesCount, _ := strconv.Atoi(strings.TrimSpace(candlesCountTmp))

	fmt.Printf("Enter %d candle's heights:\n", candlesCount)

	var candles []int32
	candlesTemp, _ := reader.ReadString('\n')
	candlesArr := strings.Split(candlesTemp, " ")

	for i := 0; i < candlesCount; i++ {
		candleTemp, _ := strconv.Atoi(strings.TrimSpace(candlesArr[i]))
		candles = append(candles, int32(candleTemp))
	}

	res := birthdayCakeCandles(candles)
	fmt.Printf("Total tallest candle(s): %d", res)
}
