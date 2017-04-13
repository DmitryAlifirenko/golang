package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"
)

var (
	hashtable = make(HashTable)
	wg        sync.WaitGroup
	dictionary_size,
	words,
	misspellings int
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func loadDictionary() {
	dictionary, err := os.Open("dictionary.txt")
	checkError(err)
	defer dictionary.Close()

	scanner := bufio.NewScanner(dictionary)

	for scanner.Scan() {
		hashtable.AddItem(scanner.Bytes())
		dictionary_size++
	}
	wg.Done()
}

func getWordFromText(ch1 chan []byte) {
	text, err := os.Open("text.txt")
	checkError(err)
	defer text.Close()

	separator := func(c rune) bool {
		return !unicode.IsLetter(c) && c != '\''
	}

	scanner := bufio.NewScanner(text)

	for scanner.Scan() {
		line := scanner.Text()
		// Split the line
		parts := strings.FieldsFunc(line, separator)
		// Loop over the parts from the string
		for i := range parts {
			words++
			ch1 <- []byte(parts[i])
		}
	}
	wg.Done()
}

func checkWord(ch1 chan []byte, ch2 chan []byte) {
	for  {
		word := <-ch1
		length := len(word)
		check_word := make([]byte, length)

		check_word = bytes.ToLower(word)

		if !hashtable.FindItem(check_word) {
			misspellings++
			ch2 <- word
		}
	}
}

func printMisspelledWords(ch chan []byte) {
	fmt.Println("Misspelled words: ")
	for  {
		fmt.Printf("%s\n", <-ch)
	}
}

func printStatistic() {
	fmt.Printf("Words in dictionary: %d\n", dictionary_size)
	fmt.Printf("Words in text: %d\n", words)
	fmt.Printf("Words misspelled: %d\n", misspellings)
}

func main() {

	Words := make(chan []byte)
	MisspelledWords := make(chan []byte)

	wg.Add(1)
	go loadDictionary()
	wg.Wait()
	wg.Add(1)
	go getWordFromText(Words)
	go checkWord(Words, MisspelledWords)
	go printMisspelledWords(MisspelledWords)
	wg.Wait()
	printStatistic()
}
