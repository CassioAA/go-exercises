package main

func RemoveAdjacentDuplicates(word string) string {

	var stack []rune

	for _, character := range word {

		length := len(stack)

		// short-circuit evaluation / minimal evaluation / McCarthy evaluation
		if length > 0 && character == stack[length-1] {
			// pop
			// if any duplicates appear in the stack it get removed
			stack = stack[:length-1]
		} else {
			// if it's not on top, it was already checked
			// push
			stack = append(stack, character)
		}

	}

	return string(stack)

}
