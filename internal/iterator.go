package internal

import (
	"fmt"
	"strings"
)

// Repeat : repeats a letter
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated = repeated + character
	}
	return repeated
}

// TrimSuffixOfString : Trims suffix
func TrimSuffixOfString(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

func workOutTypesWithOnlyStrings(foo, bar string) string {
	return foo + bar
}

func workOutTypesWithMix(foo int, bar string) string {
	return fmt.Sprintf("%d is a number, %s is a string", foo, bar)
}
