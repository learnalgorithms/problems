package main

import "fmt"

func main() {
	fmt.Println(reverseString("vim-go"))
}

func reverseString(str string) string {
	strLen := len(str)
	strByte := make([]rune, strLen)
	for idx, r := range str {
		strByte[strLen-1-idx] = r
	}
	return string(strByte)
}
