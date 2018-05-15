package main

import "fmt"

const helloPrefix = "Hello, "

func Hello() string {
	return helloPrefix + "world"
}

func HelloString(name string) string {
	if name == "" {
		name = "world"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello())
}
