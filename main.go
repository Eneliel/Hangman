package main

import (
	"fmt"
	"io/ioutil"
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
	guessedLetters := InitWords(targetWord)
	HmState := 6
	GameStages(targetWord, guessedLetters, HmState)
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

func GameStages(targetword string, guessLetters map[rune]bool, HmState int) {
	fmt.Println(GetWordGuessProgress(targetword, guessLetters))
	fmt.Println()
	fmt.Println(PrintHmStages(HmState))
}

func GetWordGuessProgress(targetWord string, guessedLetters map[rune]bool) string {
	result := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] == true {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}
		result += " "
	}
	fmt.Println("")
	return result
}

func PrintHmStages(HmState int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("Stages/Hm%d", HmState))
	if err != nil {
		panic(err)
	}

	return string(data)
}
