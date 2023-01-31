package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

var dict = []string{
	"Gopher",
	"Apple",
	"Gamer",
	"Programm",
	"Zombie",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	targetWord := dict[rand.Intn(len(dict))]
	targetWord = "Apple"
	guessedLetters := InitWords(targetWord)
	GameStages(targetWord, guessedLetters)
	guessedLetters['p'] = true
	GameStages(targetWord, guessedLetters)
	guessedLetters['l'] = true
	GameStages(targetWord, guessedLetters)
}

func RandWord() string {
	word := dict[rand.Intn(len(dict))]
	return word
}

func InitWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true
	return guessedLetters
}

func GameStages(targetword string, guessLetters map[rune]bool) {
	for _, ch := range targetword {
		if ch == ' ' {
			fmt.Print(" ")
		} else if guessLetters[unicode.ToLower(ch)] == true {
			fmt.Printf("%c", ch)
		} else {
			fmt.Print("_")
		}
		fmt.Print(" ")
	}
	fmt.Println(" ")
}
