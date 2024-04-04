package main

import "testing"

func TestRemoveAdjacentDuplicates(t *testing.T) {

	removeAdjacentDuplicates := RemoveAdjacentDuplicates("abbaca")
	expected := "ca"
	if removeAdjacentDuplicates != expected {
		t.Errorf("expected '%q' but got '%q'", 
			expected, removeAdjacentDuplicates,
		)
	}

	removeAdjacentDuplicates = RemoveAdjacentDuplicates("azxxzy")
	expected = "ay"
	if removeAdjacentDuplicates != expected {
		t.Errorf("expected '%q' but got '%q'", 
			 expected, removeAdjacentDuplicates)
	}

	removeAdjacentDuplicates = RemoveAdjacentDuplicates("abccba")
	expected = ""
	if removeAdjacentDuplicates != expected {
		t.Errorf("expected '%q' but got '%q'", 
			 expected, removeAdjacentDuplicates)
	}

	removeAdjacentDuplicates = RemoveAdjacentDuplicates("foobar")
	expected = "fbar"
	if removeAdjacentDuplicates != expected {
		t.Errorf("expected '%q' but got '%q'", 
			 expected, removeAdjacentDuplicates)
	}

	removeAdjacentDuplicates = RemoveAdjacentDuplicates("abcd")
	expected = "abcd"
	if removeAdjacentDuplicates != expected {
		t.Errorf("expected '%q' but got '%q'", 
			 expected, removeAdjacentDuplicates)
	}

}
