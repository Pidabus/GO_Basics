package main

import "fmt"

func Hello() string {
	return fmt.Sprint("Hello, world")
}

func main() {
	fmt.Println(Hello())
}