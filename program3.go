package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "abaabcdefcabad"

	result := strings.Count(str1, "a")

	fmt.Printf("a occurs %d times", result)

}
