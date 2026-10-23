package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	malay   = "Malay"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	malayHelloPrefix   = "Apa Khabar Bossku, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case malay:
		prefix = malayHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
