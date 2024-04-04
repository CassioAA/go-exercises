package main

import "testing"

func TestGetLongestCommonPrefix(t *testing.T) {

	longestCommonPrefix := GetLongestCommonPrefix([]string{"flower", "flow", "flight"})
	expected := "fl"
	if longestCommonPrefix != expected {
		t.Errorf("expected '%q' but got '%q'",
			expected, longestCommonPrefix,
		)
	}

	longestCommonPrefix = GetLongestCommonPrefix([]string{"dog", "racecar", "car"})
	expected = ""
	if longestCommonPrefix != expected {
		t.Errorf("expected '%q' but got '%q'",
			expected, longestCommonPrefix)
	}

	longestCommonPrefix = GetLongestCommonPrefix([]string{"digital", "dig", "dinner"})
	expected = "di"
	if longestCommonPrefix != expected {
		t.Errorf("expected '%q' but got '%q'",
			expected, longestCommonPrefix)
	}

	longestCommonPrefix = GetLongestCommonPrefix([]string{"happy", "harbor", "hard"})
	expected = "ha"
	if longestCommonPrefix != expected {
		t.Errorf("expected '%q' but got '%q'",
			expected, longestCommonPrefix)
	}

	longestCommonPrefix = GetLongestCommonPrefix([]string{"apple", "application", "appetite", "app"})
	expected = "app"
	if longestCommonPrefix != expected {
		t.Errorf("expected '%q' but got '%q'",
			expected, longestCommonPrefix)
	}

}
