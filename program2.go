package main

func main() {

	in := []rune{'a', 'b', 'a', 'a', 'b', 'c', 'd', 'e', 'f', 'c', 'a', 'b', 'a'}
	Map := make(map[rune]int)

	for _, v := range in {
		_, val := Map[v]
		if val {
			Map[rune(v)]++
		} else {
			Map[rune(v)] = 1
		}

		count := 1
		for i, j := range Map {
			if i > rune(count) {
				count = j
			}

		}

	}
}
