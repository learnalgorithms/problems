package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("Missing argument: You need to pass an integer argument")
	}
	for argIdx, number := range args[1:] {
		n, err := strconv.Atoi(number)
		if err != nil {
			log.Fatalf("Invalid argument: The argument needs to be an integer, at %d", argIdx)
		}
		FizzBuzz(n)
	}
}

func FizzBuzz(number int) {
	if number%3 == 0 {
		fmt.Print("Fizz")
	}
	if number%5 == 0 {
		fmt.Print("Buzz")
	}
	if number%3 != 0 && number%5 != 0 {
		fmt.Print(number)
	}
	fmt.Println()
}
