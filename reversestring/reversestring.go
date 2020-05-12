package main

import "fmt"

func main() {
	fmt.Println(reverseString("vim-go"))
}

func reverseString(str string) string {
	strLen := len(str)
	strByte := make([]byte, strLen)
	for idx, r := range str {
		strByte[strLen-1-idx] = byte(r)
	}
	return string(strByte)
}
