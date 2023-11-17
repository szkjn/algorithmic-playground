/*
Given a time in 12-hour AM/PM format, convert it to 24-hour format.

- 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
- 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.

Example:
- s = 12:01:00PM -> returns 12:01:00
- s = 12:01:00AM -> returns 00:01:00

Function parameter:
- string s: time in 12 hour format

Returns:
- string: time in 24 hour format

Input Format:
- A single string of time in 12-hour format (i.e.: hh:mm:ssAM or hh:mm:ssPM).

Constraints:
- All input times are valid
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	res := s[0:8]
	hh := s[:2]
	meridiem := strings.TrimSpace(s[8:])

	hh_tmp, err := strconv.Atoi(hh)
	if err != nil {
		panic(err)
	}

	if meridiem == "PM" && hh_tmp < 12 {
		res = strconv.Itoa(hh_tmp+12) + res[2:]
	} else if hh_tmp == 12 {
		res = "00" + res[2:]
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	fmt.Println("Enter time in 12-hours format hh:mm:ss<AM/PM>:")
	input, _ := reader.ReadString('\n')

	res := timeConversion(input)
	fmt.Println(res)
}
