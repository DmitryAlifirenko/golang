// Vigenere project vigenere.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var keycode_length int

func getText() string {
	fmt.Println("plaintext: ")

	text := bufio.NewScanner(os.Stdin)
	text.Scan()

	return text.Text()
}

func getKey() string {
	key := os.Args[1]
	return key
}

func setKeycode(key string) []rune {
	keycode_length = len(key)
	p := make([]rune, keycode_length)
	for i := 0; i < keycode_length; i++ {
		p[i] = unicode.ToUpper(rune(key[i])) - 65
	}
	return p
}

func encryptText(keycode []rune, text string) {
	key_count := 0

	length := len(text)

	fmt.Printf("ciphertext: ")

	for i := 0; i < length; i++ {
		if !unicode.IsLetter(rune(text[i])) {
			fmt.Printf("%c", text[i])
		} else {
			fmt.Printf("%c", caesar(rune(text[i]), keycode[key_count]))
			if key_count < keycode_length-1 {
				key_count++
			} else {
				key_count = 0
			}
		}
	}
}

func caesar(text1 rune, key1 rune) rune {
	firstLowercaseLetter := rune(97)
	firstUppercaseLetter := rune(65)
	alphabetLength := rune(26)

	if unicode.IsLower(text1) {
		return ((((text1 - firstLowercaseLetter) + key1) % alphabetLength) + firstLowercaseLetter)
	} else {
		return ((((text1 - firstUppercaseLetter) + key1) % alphabetLength) + firstUppercaseLetter)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage vigenere <key>\n")
		os.Exit(0)
	}

	encryptText(setKeycode(getKey()), getText())
}
