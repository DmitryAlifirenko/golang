package main

import (
	"os"
	"bufio"
	"strings"
	"unicode"
	"fmt"
	"bytes"
	"sync"
	"time"
)

var (
	hashtable = make(HashTable)
  	mutex = &sync.Mutex{}
	dictionary_size,
	words,
	misspellings int
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func  LoadDictionary()  {
	dictionary, err := os.Open("dictionary.txt")
	checkError(err)
	defer dictionary.Close()

	scanner := bufio.NewScanner(dictionary)
	mutex.Lock()
	for scanner.Scan() {
		hashtable.AddItem(scanner.Bytes())
		dictionary_size++
	}
	mutex.Unlock()
}

func  GetWordFromText(ch chan []byte)  {
	text, err := os.Open("text.txt")
	checkError(err)
	defer text.Close()

	separator := func(c rune) bool {
		return  !unicode.IsLetter(c) && c!='\''
	}

	scanner := bufio.NewScanner(text)

	for scanner.Scan() {
		line := scanner.Text()
		// Split the line
		parts := strings.FieldsFunc(line, separator)
		// Loop over the parts from the string
		for i := range parts {
			words++
			ch <- []byte(parts[i])
		}
	}
}


func  CheckWord(ch chan []byte, ch1 chan []byte)  {
	for <-ch != nil{
		word := <-ch
		length:= len(word)
		check_word:=make([]byte, length)

		check_word = bytes.ToLower(word)
		mutex.Lock()
		if !hashtable.FindItem(check_word) {
			misspellings++
			ch1<-word
		}
		mutex.Unlock()
	}
}

func  PrintMisspelledWords(ch chan []byte) {
	fmt.Println("Misspelled words: ")
	for  {
		fmt.Printf("%s\n", <-ch)
	}
}

func PrintStatistic() {
	time.Sleep(time.Second)
	fmt.Printf("Words in dictionary: %d\n", dictionary_size)
	fmt.Printf("Words in text: %d\n", words)
	fmt.Printf("Words misspelled: %d\n", misspellings)
}

func main() {

	Words := make(chan []byte)
	MisspelledWords := make(chan []byte)

	go LoadDictionary()
	go GetWordFromText(Words)
	go CheckWord(Words, MisspelledWords)
	go PrintMisspelledWords(MisspelledWords)
	go PrintStatistic()

	time.Sleep(time.Second*2)
}