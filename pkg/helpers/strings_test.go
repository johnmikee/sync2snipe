package helpers

import (
	"errors"
	"testing"
)

func TestCompareOSVer(t *testing.T) {
	testCases := []struct {
		a              string
		b              string
		expectedResult bool
		expectedError  error
	}{
		{"1.2.3", "1.2.4", false, nil},
		{"1.2.3", "1.2.3", true, nil},
		{"1.2.4", "1.2.3", true, nil},
		{"1.2.3", "2.0.0", false, nil},
		{"invalid", "1.2.3", false, InvalidSemVerError{"Invalid semantic version"}},
	}

	for _, testCase := range testCases {
		result, err := CompareOSVer(testCase.a, testCase.b)

		if result != testCase.expectedResult {
			t.Errorf("Test case failed: a=%s, b=%s, expected result=%v, got=%v", testCase.a, testCase.b, testCase.expectedResult, result)
		}

		if !errors.Is(err, testCase.expectedError) {
			t.Errorf("Test case failed: a=%s, b=%s, expected error=%v, got=%v", testCase.a, testCase.b, testCase.expectedError, err)
		}
	}
}

func TestContains(t *testing.T) {
	slice := []string{"apple", "banana", "cherry", "date", "elderberry"}

	// Test case: element exists in the slice
	exists := "banana"
	result := Contains(slice, exists)
	if !result {
		t.Errorf("Test case failed: expected %s to be found in slice", exists)
	}

	// Test case: element does not exist in the slice
	notExists := "grape"
	result = Contains(slice, notExists)
	if result {
		t.Errorf("Test case failed: unexpected %s found in slice", notExists)
	}
}

func TestContainsPosition(t *testing.T) {
	slice := []string{"apple", "banana", "cherry", "date", "elderberry"}

	// Test case: element exists in the slice
	exists := "cherry"
	expectedIndex := 2
	resultIndex, result := ContainsPosition(slice, exists)
	if !result {
		t.Errorf("Test case failed: expected %s to be found in slice", exists)
	}
	if resultIndex != expectedIndex {
		t.Errorf("Test case failed: expected index=%d, got index=%d", expectedIndex, resultIndex)
	}

	// Test case: element does not exist in the slice
	notExists := "grape"
	expectedIndex = 0
	resultIndex, result = ContainsPosition(slice, notExists)
	if result {
		t.Errorf("Test case failed: unexpected %s found in slice", notExists)
	}
	if resultIndex != expectedIndex {
		t.Errorf("Test case failed: expected index=%d, got index=%d", expectedIndex, resultIndex)
	}
}

func TestPossessiveForm(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult string
	}{
		{"John", "John's"},
		{"Mary", "Mary's"},
		{"Chris", "Chris'"},
		{"James", "James'"},
		{"Lucy", "Lucy's"},
		{"", ""},
	}

	for _, testCase := range testCases {
		result := PossessiveForm(testCase.name)
		if result != testCase.expectedResult {
			t.Errorf("Test case failed: name=%s, expected=%s, got=%s", testCase.name, testCase.expectedResult, result)
		}
	}
}

func TestRemove(t *testing.T) {
	slice := []string{"apple", "banana", "cherry", "date", "elderberry"}
	index := 2 // Index of "cherry" to be removed
	expectedResult := []string{"apple", "banana", "date", "elderberry"}

	result := Remove(slice, index)

	if len(result) != len(expectedResult) {
		t.Errorf("Test case failed: expected length=%d, got=%d", len(expectedResult), len(result))
	}

	for i := range result {
		if result[i] != expectedResult[i] {
			t.Errorf("Test case failed: expected %s, got %s", expectedResult[i], result[i])
		}
	}
}
