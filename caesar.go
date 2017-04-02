// Caesar project caesar.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func getText() string {
	fmt.Println("plaintext: ")

	text := bufio.NewScanner(os.Stdin)
	text.Scan()

	return text.Text()
}

func encryptText(text string) {
	firstLowercaseLetter := byte(97)
	firstUppercaseLetter := byte(65)
	alphabetLength := byte(26)
	length := len(text)

	key, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("ciphertext: ")

	for i := 0; i < length; i++ {
		if unicode.IsLetter(rune(text[i])) {
			if unicode.IsLower(rune(text[i])) {
				fmt.Printf("%c", ((((text[i] - firstLowercaseLetter) + byte(key)) % alphabetLength) + firstLowercaseLetter))
			} else {
				fmt.Printf("%c", ((((text[i] - firstUppercaseLetter) + byte(key)) % alphabetLength) + firstUppercaseLetter))
			}
		} else {
			fmt.Printf("%c", text[i])
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage caesar <key>\n")
		os.Exit(0)
	}

	encryptText(getText())
}
