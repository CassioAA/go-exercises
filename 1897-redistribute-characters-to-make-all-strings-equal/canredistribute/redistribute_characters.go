package main

func CanMakeEqual(words []string) bool {

	// var charCounts map[rune]int doesn't allocate space in memory
	// both of the below allocate
	// charCounts := map[rune]int{}
	charCounts := make(map[rune]int)

	// the indexes are not used
	for _, word := range words {
		for _, character := range word {
			// if the character is a new key, it creates a new one
			charCounts[character]++
		}
	}

	// if each character's frequency is divisible by the number of strings, they can be equally redistributed
	numWords := len(words)
	for _, charCount := range charCounts {
		if charCount%numWords != 0 {
			return false
		}
	}

	return true

}