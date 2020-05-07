package main

import "fmt"

func main() {
	fmt.Println(hasUniqueChars("repeated characters"))
	fmt.Println(hasUniqueCharsNoAuxStruct("no reps"))
}

func hasUniqueChars(str string) bool {
	chars := make([]int, 256)

	if len(str) > 256 {
		return false
	}

	for _, char := range str {
		asciiCode := int(char)
		chars[asciiCode] += 1
		if chars[asciiCode] > 1 {
			return false
		}
	}
	return true
}

func hasUniqueCharsNoAuxStruct(str string) bool {
	for idx, char1 := range str {
		for _, char2 := range str[idx+1:] {
			if char1 == char2 {
				return false
			}
		}
	}
	return true
}
