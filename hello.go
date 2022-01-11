package main

import "fmt"

const prefixHello = "Hello"

func Hello(greet string) string {
	if greet == ""{
		greet = "!"
	}

	return prefixHello + " " + greet
	
}

func main() {
	fmt.Println(Hello("World"))
}