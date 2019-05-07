package main

import (
	"fmt"
	"strings"
)

func main() {
	toUpperSync("Hello Callbacks!", func(v string) {
		fmt.Printf("Callback Function: %s\n", v)
	})
}

func toUpperSync(word string, f func(string)) {
	fmt.Println("Function!")
	f(strings.ToUpper(word))
}