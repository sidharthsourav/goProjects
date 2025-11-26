package stringutils

import "fmt"
import "strings"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println("in the reverse function")
	return string(runes)
}

func StringReverse(message string) string {
	chars := strings.Split(message, "")
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return strings.Join(chars, "")
}
func reverse(number int32) bool {
	fmt.Println(number)
	return true
}
