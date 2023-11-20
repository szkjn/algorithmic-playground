/*
Given n names and phone numbers, assemble a phone book that maps friends' names
to their respective phone numbers. You will then be given an unknown number of
names to query your phone book for. For each name queried, print the associated
entry from your phone book on a new line in the form name=phoneNumber; if an entry
for names is not found, print "Not found" instead.

Note: Your phone book should be a Dictionary/Map/HashMap data structure.

Input Format
- The first line contains an integer, n, denoting the number of entries in the
phone book. Each of the n subsequent lines describes an entry in the form of
2 space-separated values on a single line. The first value is a friend's name,
and the second value is an 8-digit phone number. After the n lines of phone book
entries, there are an unknown number of lines of queries. Each line (query)
contains a name to look up, and you must continue reading lines until there is
no more input.

Note: Names consist of lowercase English alphabetic letters and are first names only.

Output Format:
- On a new line for each query, print "Not found" if the name has no corresponding
entry in the phone book; otherwise, print the full name and phoneNumber in the
format name=phoneNumber.

Example:

Input:

	3
	sam 99912222
	tom 11122222
	harry 12299933
	sam
	edward
	harry

Output:

	sam=99912222
	Not found
	harry=12299933
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createPhoneBook(arr []string) map[string]string {
	res := make(map[string]string)
	for _, el := range arr {
		tmpArr := strings.Split(el, " ")
		name := strings.TrimSpace(tmpArr[0])
		phone := strings.TrimSpace(tmpArr[1])
		res[name] = phone
	}
	return res
}

func getPhone(n string, phoneBook map[string]string) string {
	res := ""
	if _, ok := phoneBook[n]; ok {
		res = n + "=" + phoneBook[n]
	} else {
		res = "Not found"
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTmp, err := reader.ReadString('\n')
	must(err)

	n, err := strconv.Atoi(strings.TrimSpace(nTmp))
	must(err)

	arr := []string{}
	for i := 0; i < n; i++ {
		contact, err := reader.ReadString('\n')
		must(err)
		arr = append(arr, contact)
	}
	fmt.Println("arr", arr)
	phoneBook := createPhoneBook(arr)
	fmt.Println("phoneBook", phoneBook)

	for {
		newContact, _ := reader.ReadString('\n')
		newContact = strings.TrimSpace(newContact)
		if newContact == "" {
			break
		}
		fmt.Println(getPhone(newContact, phoneBook))
	}

}

func must(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
