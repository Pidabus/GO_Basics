package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	answerArray := [][]string{}
	
	for key, _ := range m {
		answerArray = append(answerArray, m[key])
	}

	return answerArray
}

func main() {
	string01 := []string{"cat", "atc", "cta"}

	fmt.Println(groupAnagrams(string01))
}