package main

import "fmt"

const helloPrefixUS = "Hello, "
const helloPrefixSpanish = "Hola, "
const helloPrefixIndia = "Namaste, "

const spanish = "Spanish"
const india = "India"

func Hello() string {
	return helloPrefixUS + "world"
}

func HelloString(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = helloPrefixSpanish

	case india:
		prefix = helloPrefixIndia

	default:
		prefix = helloPrefixUS
	}

	return
}

func main() {
	fmt.Println(Hello())
}
