package main

import "fmt"

const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris", "Spanish"))
}
