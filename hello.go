package main

import "fmt"

const helloPrefix = "Hello, "
const helloPrefixSpanish = "Hola, "
const helloPrefixIndia = "Namaste, "

const spanish = "Spanish"
const india = "India"

func Hello() string {
	return helloPrefix + "world"
}

func HelloString(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := helloPrefix

	if language == spanish {
		prefix = helloPrefixSpanish
	}

	if language == india {
		prefix = helloPrefixIndia
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello())
}
