/*
Code Wars Kata :
Your task is to implement a function that would take the morse code as input and return a decoded human-readable string.
*/

package main

import (
	"fmt"
	"strings"
)

var morseCode = map[string]string{
	".-": "A", "-...": "B", "-.-.": "C", "-..": "D", ".": "E", "..-.": "F",
	"--.": "G", "....": "H", "..": "I", ".---": "J", "-.-": "K", ".-..": "L",
	"--": "M", "-.": "N", "---": "O", ".--.": "P", "--.-": "Q", ".-.": "R",
	"...": "S", "-": "T", "..-": "U", "...-": "V", ".--": "W", "-..-": "X",
	"-.--": "Y", "--..": "Z",
	"-----": "0", ".----": "1", "..---": "2", "...--": "3", "....-": "4",
	".....": "5", "-....": "6", "--...": "7", "---..": "8", "----.": "9",
	".-.-.-": ".", "--..--": ",", "..--..": "?", ".----.": "'", "-.-.--": "!",
	"-..-.": "/", "-.--.": "(", "-.--.-": ")", ".-...": "&", "---...": ":",
	"-.-.-.": ";", "-...-": "=", ".-.-.": "+", "-....-": "-", "..--.-": "_",
	".-..-.": "\"", "...-..-": "$", ".--.-.": "@", "...---...": "SOS",
}

func decodeMorse(morseCodeStr string) string {
	words := strings.Split(strings.TrimSpace(morseCodeStr), "   ")
	decodedWords := make([]string, len(words))

	for i, word := range words {
		letters := strings.Split(word, " ")
		decodedLetters := make([]string, len(letters))

		for j, letter := range letters {
			if char, ok := morseCode[letter]; ok {
				decodedLetters[j] = char
			}
		}
		decodedWords[i] = strings.Join(decodedLetters, "")
	}
	return strings.Join(decodedWords, " ")
}

func main() {
	fmt.Println(decodeMorse(".... . -.--   .--- ..- -.. .")) // Output: 'HEY JUDE'
	fmt.Println(decodeMorse("...---..."))                    // Output: 'SOS'
	fmt.Println(decodeMorse("...   ---   ..."))              // Output: 'S O S'
}
