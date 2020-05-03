package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	var num []int
	if len(args) < 2 {
		log.Fatal("Missing argument, you need to pass a number as argument")
	}
	for _, v := range args[1:] {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Invalid argument, you need to pass a number as argument")
		}
		num = append(num, n)
	}
	FizzBuzz(num...)
}

func FizzBuzz(number ...int) {
	for _, n := range number {
		if n%3 == 0 {
			fmt.Print("Fizz")
		}
		if n%5 == 0 {
			fmt.Print("Buzz")
		}
		if n%5 != 0 && n%3 != 0 {
			fmt.Print(n)
		}
		fmt.Println()
	}
}
