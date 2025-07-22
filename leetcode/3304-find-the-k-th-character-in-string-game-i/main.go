package main

import "fmt"

func kthCharacter(k int) byte {
	words := []byte{'a'}
	for len(words) < k {
		for _, char := range words {
			words = append(words, 'a'+(char-'a'+1)%26)
		}
	}

	fmt.Println(words)
	return words[k-1]
}

func main() {
	k := 10
	result := kthCharacter(k)
	fmt.Println(byte('z' - 'a' + 1))
	fmt.Printf("Result when k = %d: %c\n", k, result)
}
