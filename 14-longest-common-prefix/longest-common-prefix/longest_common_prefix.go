package main

func GetLongestCommonPrefix(words []string) string {

    if len(words) == 0 {
        return ""
    }

	// assuming the entire first word is a prefix
    prefix := words[0]
	
	// iterates starting from the second word
    for _, word := range words[1:] {
		// while the prefix doesn't match the start of the current string
        for len(word) < len(prefix) || word[:len(prefix)] != prefix {
            prefix = prefix[ : len(prefix)-1 ]
			// if the prefix becomes empty, returns empty word straight away
            if prefix == "" {
                return ""
            }
        }
    }

    return prefix

}
