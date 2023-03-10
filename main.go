package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var ReadInput = bufio.NewReader(os.Stdin)
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
	HmState := 0
	for !isGameover(targetWord, HmState, guessedLetters) {
		GameStages(targetWord, guessedLetters, HmState)
		input := readInput()
		if len(input) != 1 {
			fmt.Println("Please, use only 1 letter")
			continue
		}
		letter := rune(input[0])
		if isCorrectGuess(targetWord, letter) {
			guessedLetters[letter] = true
		} else {
			HmState++
		}
	}
	fmt.Println("Game over...")
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("You Win!")
	} else if isHmComplete(HmState) {
		fmt.Println("Tou Lose!")
	} else {
		panic("invalid state!")
	}
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

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, ch := range targetWord {
		if !guessedLetters[unicode.ToLower(ch)] {
			return false
		}
	}

	return true
}

func isHmComplete(HmState int) bool {
	return HmState >= 9
}

func GameStages(targetword string, guessLetters map[rune]bool, HmState int) {
	fmt.Println(GetWordGuessProgress(targetword, guessLetters))
	fmt.Println()
	fmt.Println(PrintHmStages(HmState))
}

func isGameover(targetWord string, HmState int, guessedWord map[rune]bool) bool {
	return isWordGuessed(targetWord, guessedWord) || isHmComplete(HmState)
}

func GetWordGuessProgress(targetWord string, guessedLetters map[rune]bool) string {
	result := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] {
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

func readInput() string {
	fmt.Print("> ")
	input, err := ReadInput.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

func isCorrectGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)
}
