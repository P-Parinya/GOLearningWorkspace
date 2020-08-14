package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	counter := map[string]int{}

	//break long string into each word.
	wordList := strings.Split(s," ")

	//loop check word by word.
	for _,word:= range wordList{
		
		word = strings.TrimSuffix(word,",")
		
		_ , ok := counter[word] //check that word existed.
		if ok {
			//if which word was existed, increase count number by 1.
			counter[word] += 1
		}else{
			//if which word was not existed, initial count number.
			counter[word] = 1
		}
	}
	
	return counter
}

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck"
	w := WordCount(s)
	fmt.Println(w)
}
