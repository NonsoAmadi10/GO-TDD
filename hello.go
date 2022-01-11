package main

import "fmt"

const englishPrefixHello = "Hello, "
const spanish = "spanish"
const french = "french"
const spanishPrefixHello = "Hola, "
const frenchPrefixHello  = "Bonjour, "

func Hello(greet string, lang string) string {
	if greet == ""{
		greet = "World"
	}

	prefix := englishPrefixHello

	switch lang {
		case french:
			prefix = frenchPrefixHello
		case spanish:
			prefix = spanishPrefixHello	
	}

	return prefix  + greet
	
}

func main() {
	fmt.Println(Hello("World", "eng"))
}