package storage

import "testing"

func Test_GetUploadPath(t *testing.T) {
	// Test empty input
	emptyPath, err := GetUploadPath("")
	if emptyPath != "" || err == nil {
		t.Errorf("Expected empty path and non-nil error for empty input, got path: %s, error: %v", emptyPath, err)
	}

	// Test input with only letters and digits
	input := "Hello2You"
	expected := "He/ll/o2/Yo/u"
	result, err := GetUploadPath(input)
	if result != expected || err != nil {
		t.Errorf("Expected: %s, got: %s, error: %v", expected, result, err)
	}

	// Test input with special characters
	input = "Goodbye 3 W-o-r-l-d?"
	expected = "Go/od/by/e3/Wo/rl/d"
	result, err = GetUploadPath(input)
	if result != expected || err != nil {
		t.Errorf("Expected: %s, got: %s, error: %v", expected, result, err)
	}

	// Test even number of characters
	input = "HelloYou"
	expected = "He/ll/oY/ou"
	result, err = GetUploadPath(input)
	if result != expected || err != nil {
		t.Errorf("Expected: %s, got: %s, error: %v", expected, result, err)
	}
}
