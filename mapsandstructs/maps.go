package mapsandstructs

import (
	"fmt"
)

//MapPlay1 will make a testmap and print it then do some training!.
func MapPlay1() {
	wordCount := make(map[string]int)
	wordCount = map[string]int{
		"Hello": 1,
		"Bye":   1,
		"The":   12,
		"Lol":   7,
	}
	fmt.Println(wordCount)
	wordCount["Fiskepudding"] = 99
	delete(wordCount, "Fiskepudding")

	fmt.Println(wordCount["Fiskepudding"])
	//Notice: No error raised even though key does not exist!

	count, ok := wordCount["Fiskepudding"]
	fmt.Println(count, ok)
	//This count, ok can be used to verify wether the key exists

	wc := wordCount
	delete(wc, "Hello")
	fmt.Println(wc)
	fmt.Println(wordCount)
	//Shallow copy, change in wc changes the WordCount too
}
