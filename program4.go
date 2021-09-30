package main

import "fmt"

func main() {
	str1 := "abaabcdefcabad"

	count := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] == 'a' {
			count++
		}

	}
	fmt.Println(count)

}
