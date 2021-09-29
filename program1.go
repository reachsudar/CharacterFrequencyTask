package main

import "fmt"

func main() {

	a := []rune{'a', 'b', 'a', 'a', 'b', 'c', 'd', 'e', 'f', 'c', 'a', 'b', 'a', 'd'}

	count := 0
	for _, v := range a {
		if v == 'a' {
			count++
		}
	}
	fmt.Printf("a occurs %d times", count)

}
