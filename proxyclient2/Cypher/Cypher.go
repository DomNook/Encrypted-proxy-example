package Cypher

import (
	"fmt"
)

var alphabet = "abcdefghijklmnopqrstuvwxyzæøå,. 0123456789"
var alphArray = []rune(alphabet)

func EncryptLetter(toEncrypt rune) rune {
	letterPos := 0
	var found = false
	for index := 0; !found; index++ {
		if alphArray[index] == toEncrypt {
			found = true
			letterPos = index
		}
	}

	letterPos = letterPos + 4

	if letterPos > 41 {
		letterPos = letterPos - 42
	}
	letters := []rune(alphabet)
	fmt.Println(string(toEncrypt) + " -> " + string(letters[letterPos]))
	return letters[letterPos]

}

func DecryptLetter(toDecrypt rune) rune {
	letterPos := 0
	var found = false
	for index := 0; !found; index++ {
		if alphArray[index] == toDecrypt {
			found = true
			letterPos = index
		}
	}

	letterPos = letterPos - 4

	if letterPos < 0 {
		letterPos = letterPos + 42
	}
	letters := []rune(alphabet)
	fmt.Println(string(toDecrypt) + " -> " + string(letters[letterPos]))
	return letters[letterPos]

}
