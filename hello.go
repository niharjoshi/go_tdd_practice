package main

import "fmt"

const helloPrefixEnglish = "Hello, "
const spanish = "Spanish"
const helloPrefixSpanish = "Hola, "
const french = "French"
const helloPrefixFrench = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	prefix = helloPrefixEnglish
	switch language {
	default:
		prefix = helloPrefixEnglish
	case spanish:
		prefix = helloPrefixSpanish
	case french:
		prefix = helloPrefixFrench
	}
	return
}

func main() {
	fmt.Println(Hello("Nihar", ""))
}
