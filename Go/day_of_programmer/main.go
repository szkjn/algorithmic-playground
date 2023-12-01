package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func dayOfProgrammer(year int32) string {
	if year == 1918 {
		// Russian calendar switch
		return "26.09.1918"
	}

	leapYear := false
	if year < 1918 {
		// Julian calendar
		leapYear = year%4 == 0
	} else {
		// Gregorian calendar
		leapYear = (year%400 == 0) || (year%4 == 0 && year%100 != 0)
	}

	if leapYear {
		return fmt.Sprintf("12.09.%d", year)
	}
	return fmt.Sprintf("13.09.%d", year)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := dayOfProgrammer(year)

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
