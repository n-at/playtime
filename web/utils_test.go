package web

import "testing"

func Test_cleanupName(t *testing.T) {
	// Test empty input
	emptyName := ""
	expectedEmptyResult := ""
	emptyResult := cleanupName(emptyName)
	if emptyResult != expectedEmptyResult {
		t.Errorf("Expected: %s, got: %s", expectedEmptyResult, emptyResult)
	}

	// Test input without underscores or dots
	noUnderscoreOrDot := "filename"
	expectedNoUnderscoreOrDotResult := "filename"
	noUnderscoreOrDotResult := cleanupName(noUnderscoreOrDot)
	if noUnderscoreOrDotResult != expectedNoUnderscoreOrDotResult {
		t.Errorf("Expected: %s, got: %s", expectedNoUnderscoreOrDotResult, noUnderscoreOrDotResult)
	}

	// Test input with underscores
	withUnderscores := "hello_world.txt"
	expectedWithUnderscoresResult := "hello world"
	withUnderscoresResult := cleanupName(withUnderscores)
	if withUnderscoresResult != expectedWithUnderscoresResult {
		t.Errorf("Expected: %s, got: %s", expectedWithUnderscoresResult, withUnderscoresResult)
	}

	// Test input with underscores and dots
	withUnderscoresAndDots := "good_file_name.docx"
	expectedWithUnderscoresAndDotsResult := "good file name"
	withUnderscoresAndDotsResult := cleanupName(withUnderscoresAndDots)
	if withUnderscoresAndDotsResult != expectedWithUnderscoresAndDotsResult {
		t.Errorf("Expected: %s, got: %s", expectedWithUnderscoresAndDotsResult, withUnderscoresAndDotsResult)
	}
}
