package main

import "testing"

func TestCanMakeEqual(t *testing.T) {

	canMakeEqual := CanMakeEqual([]string{"aaa", "bbb", "ccc", "abc"})
	expected := true
	if canMakeEqual != expected {
		t.Errorf("expected '%t' but got '%t'", expected, canMakeEqual)
	}

	canMakeEqual = CanMakeEqual([]string{"ab", "bc", "cd", "de", "ea"})
	expected = false
	if canMakeEqual != expected {
		t.Errorf("expected '%t' but got '%t'", expected, canMakeEqual)
	}

	canMakeEqual = CanMakeEqual([]string{"zzx", "xzz", "zxz"})
	expected = true
	if canMakeEqual != expected {
		t.Errorf("expected '%t' but got '%t'", expected, canMakeEqual)
	}

	canMakeEqual = CanMakeEqual([]string{"abc", "aabc", "bc"})
	expected = true
	if canMakeEqual != expected {
		t.Errorf("expected '%t' but got '%t'", expected, canMakeEqual)
	}

	canMakeEqual = CanMakeEqual([]string{"ab", "a"})
	expected = false
	if canMakeEqual != expected {
		t.Errorf("expected '%t' but got '%t'", expected, canMakeEqual)
	}

}
