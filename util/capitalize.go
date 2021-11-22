package util

import (
	"fmt"
	"strings"
	"unicode"
)

// Capitalises the first character of a string
func Capitalise(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(strings.ToLower(str))
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}

//Error utiliy function
func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
