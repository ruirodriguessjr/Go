package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println(t)

	// Example 1: Willkommen to GoLangCode.com
	myText := "Welcome to GoLangCode.com"
	myText = strings.Replace(myText, "Welcome", "Willkommen", -1)
	fmt.Println(myText)

	// Example 2: Change first occurance
	// Output: The car sounds sound
	myText = "The sound sounds sound"
	myText = strings.Replace(myText, "sound", "car", 1)
	fmt.Println(myText)

	// Example 3: Replacing quotes (double backslash needed)
	// Output: I \'quote\' this text
	myText = "I 'quote' this text"
	myText = strings.Replace(myText, "'", "\\'", -1)
	fmt.Println(myText)
}
