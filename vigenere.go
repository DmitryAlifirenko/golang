// Vigenere project vigenere.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

const (
	ASCII_OF_a = 97
	ASCII_OF_A = 65
	ALPHABET_SIZE = 26
)

var keycode_length int

func checkCommandLineArgs() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage caesar <key>\n")
		os.Exit(0)
	}
}

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

func setText(input_text string) []rune {
	text_length := len(input_text)
	text:= make([]rune, text_length)
	for i:=0; i<text_length; i++ {
		text[i] = rune(input_text[i])
	}
	return text
}

func setKeycode(key string) []rune {
	keycode_length = len(key)
	p := make([]rune, keycode_length)
	for i := 0; i < keycode_length; i++ {
		p[i] = unicode.ToUpper(rune(key[i])) - ASCII_OF_A
	}
	return p
}

func printEncryptedText(keycode []rune, text []rune) {
	key_count := 0
	length := len(text)
	fmt.Println("ciphertext: ")

	for i := 0; i < length; i++ {
		fmt.Printf("%c", caesarEncryption(text[i], keycode[key_count]))
		if key_count < keycode_length-1 {
			key_count++
		} else {
			key_count = 0
		}
	}
}

func caesarEncryption(char rune, key rune) rune {
	if unicode.IsLetter(char) {
		if unicode.IsLower(char) {
			return ((((char - ASCII_OF_a) + key) % ALPHABET_SIZE) + ASCII_OF_a)
		} else {
			return ((((char - ASCII_OF_A) + key) % ALPHABET_SIZE) + ASCII_OF_A)
		}
	} else {
		return char
	}
}

func main() {
	checkCommandLineArgs()

	printEncryptedText(setKeycode(getKey()), setText(getText()))
}
