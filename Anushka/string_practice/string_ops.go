package stringpractice

import "unicode/utf8"

func CountState(s string) (int, int) {
	bytes := len(s)
	runes := utf8.RuneCountInString(s)
	return bytes, runes
}


func ReplaceChar(s string, newChar rune) string{
	r := []rune(s)
	r[0] = newChar
	return string(r)
}