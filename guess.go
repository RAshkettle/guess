package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("I'm thinking of a Number between 1 and 20.")

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 20
	magicNumber := rand.Intn(max-min+1) + min

	playerGuess := 0
	attempts := 0

	scanner := bufio.NewScanner(os.Stdin)

	for playerGuess != magicNumber {
		fmt.Printf("Guess the number: ")
		scanner.Scan()
		guess, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Please enter only whole numbers.")
			continue
		} else {
			playerGuess = guess
		}

		if playerGuess > magicNumber {
			fmt.Println("Too High!")
		}
		if playerGuess < magicNumber {
			fmt.Println("Too Low!")
		}

		attempts++
	}

	fmt.Printf("You won!  It only took %v attempts.\n", attempts)

	fmt.Println(magicNumber)

}
