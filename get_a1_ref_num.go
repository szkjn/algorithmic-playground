/*
Interview question of the week by Cassidy Williams (issue #286) :

Spreadsheet programs often use the A1 Reference Style to refer to
columns. Given a column name in this style, return its column number.

Examples of column names to their numbers:
A -> 1
B -> 2
C -> 3
Z -> 26
AA -> 27
AB -> 28 
AAA -> 703
*/

package main

import (
	"fmt"
)

func columnNumber(colName string) int {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	colNum := 0

	for _, c := range colName {
		index := int(c - 'A' + 1)
		colNum = colNum*len(alphabet) + index
	}

	return colNum
}

func main() {
	fmt.Println(columnNumber("A"))
	fmt.Println(columnNumber("AA"))
	fmt.Println(columnNumber("AAA"))
}
