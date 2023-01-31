package main

import (
	"fmt"
	"math/rand"
	"time"
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

	fmt.Println(targetWord)
}
