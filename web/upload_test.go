package web

import "testing"

func TestGetUploadPath(t *testing.T) {
	// Test empty input
	emptyPath, err := getUploadPath("")
	if emptyPath != "" || err == nil {
		t.Errorf("Expected empty path and non-nil error for empty input, got path: %s, error: %v", emptyPath, err)
	}

	// Test input with only letters and digits
	input := "Hello2You"
	expected := "He/ll/o2/Yo/u"
	result, err := getUploadPath(input)
	if result != expected || err != nil {
		t.Errorf("Expected: %s, got: %s, error: %v", expected, result, err)
	}

	// Test input with special characters
	input = "Goodbye 3 W-o-r-l-d?"
	expected = "Go/od/by/e3/Wo/rl/d"
	result, err = getUploadPath(input)
	if result != expected || err != nil {
		t.Errorf("Expected: %s, got: %s, error: %v", expected, result, err)
	}

	// Test even number of characters
	input = "HelloYou"
	expected = "He/ll/oY/ou"
	result, err = getUploadPath(input)
	if result != expected || err != nil {
		t.Errorf("Expected: %s, got: %s, error: %v", expected, result, err)
	}
}

func TestGetFileExtension(t *testing.T) {
	// Test empty input
	emptyName := ""
	expectedEmptyResult := ""
	emptyResult := getFileExtension(emptyName)
	if emptyResult != expectedEmptyResult {
		t.Errorf("Expected: %s, got: %s", expectedEmptyResult, emptyResult)
	}

	// Test input without an extension
	noExtensionName := "filename"
	expectedNoExtensionResult := ""
	noExtensionResult := getFileExtension(noExtensionName)
	if noExtensionResult != expectedNoExtensionResult {
		t.Errorf("Expected: %s, got: %s", expectedNoExtensionResult, noExtensionResult)
	}

	// Test input with an extension
	withExtensionName := "document.pdf"
	expectedWithExtensionResult := "pdf"
	withExtensionResult := getFileExtension(withExtensionName)
	if withExtensionResult != expectedWithExtensionResult {
		t.Errorf("Expected: %s, got: %s", expectedWithExtensionResult, withExtensionResult)
	}

	// Test input with multiple dots
	multipleDotsName := "file.name.with.dots.txt"
	expectedMultipleDotsResult := "txt"
	multipleDotsResult := getFileExtension(multipleDotsName)
	if multipleDotsResult != expectedMultipleDotsResult {
		t.Errorf("Expected: %s, got: %s", expectedMultipleDotsResult, multipleDotsResult)
	}
}
