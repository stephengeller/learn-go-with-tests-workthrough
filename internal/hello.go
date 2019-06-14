package internal

import (
	"fmt"
	"strings"
)

const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"

func greetingPrefix(language string) (prefix string) {
	language = strings.ToLower(language)
	switch language {
	case "spanish":
		prefix = spanishHelloPrefix
	case "french":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

// Hello : Says hello to you!
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)
	return fmt.Sprintf("%s, %s", prefix, name)
}

func SayHelloInLanguage() {
	fmt.Println(Hello("Stephen", "English"))
}
