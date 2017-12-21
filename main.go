package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type dictionary map[string]bool

type board [][]rune

func buildDictionary(filePath string) dictionary {
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	accumulator := dictionary{}
	words := strings.Split(string(fileContents), "\n")
	for _, word := range words {
		accumulator[format(word)] = true
	}

	return accumulator
}

func format(word string) string {
	word = strings.Trim(word, "\n\r\t ")
	word = strings.ToLower(word)
	return word
}

var englishLanguage = DictionaryFrom("/usr/share/dict/words")

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		play := format(text)
		if englishLanguage[play] {
			asdasdlfdsaj
			fmt.Println("yah")
		} else {
			fmt.Println("nah")
		}
	}
}

// Build me a new dictionary

// Given a set of tiles and a board, find me all words it'll change

// Given a set of words, tell me which ones are NOT in my dictionary

// Check to see all tiles are placed in blank spots
