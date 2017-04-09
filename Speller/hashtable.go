package main

import (
	"unicode"
)

const HASHMAP_SIZE = 100000

type Node struct {
	next  *Node
	word []byte
}

type HashTable  map[int]*Node

func hashing(key []byte) int {
	var (
		hash int
		n rune
	)
	for i:=0; i< len(key) ; i++ {

		if unicode.IsLetter(rune(key[i])) {
			n = rune(key[i]) - 'a' + 1
		} else {
			n = 27
		}
	hash = ((hash << 3) + int(n)) % HASHMAP_SIZE
	}
	return hash
}

func (hashtable HashTable) AddItem(word []byte) {
	var new_word *Node
	new_word  = new(Node)
	new_word.word = word

	index := hashing(word)

	if hashtable[index] == nil {
		hashtable[index] = new_word
		new_word.next = nil
	} else {
		new_word.next = hashtable[index]
		hashtable[index] = new_word
	}
	//fmt.Printf("%s\n",hashtable[index].word)
}

func (hashtable HashTable) FindItem(word []byte) bool {
	index := hashing(word)

	if hashtable[index] != nil {
		//fmt.Printf("%s | was found in the Hash Table!\n", word)
		return true
	}

	return false
}
