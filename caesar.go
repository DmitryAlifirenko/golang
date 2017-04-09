// Caesar project caesar.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const (
	ASCII_OF_a = 97
	ASCII_OF_A = 65
	ALPHABET_SIZE = 26
)

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

func getKey() rune {
	key, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	return rune(key)
}

func setText(input_text string) []rune {
	text_length := len(input_text)
	text:= make([]rune, text_length)
	for i:=0; i<text_length; i++ {
		text[i] = rune(input_text[i])
	}
	return text
}

func printEncryptedText(text []rune, key rune)  {
	length := len(text)
	fmt.Println("ciphertext: ")
	for i := 0; i < length; i++ {
		fmt.Printf("%c", caesarEncryption(text[i], key))
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
	printEncryptedText(setText(getText()), getKey())
}
