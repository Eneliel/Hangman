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
	for {
		GameStages(targetWord, guessedLetters, HmState)
		input := readInput()
		if len(input) != 1 {
			fmt.Println("Invalid input")
			continue
		}
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

func readInput() string {
	fmt.Print("> ")
	input, err := ReadInput.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}
