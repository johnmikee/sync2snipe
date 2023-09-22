package helpers

import (
	"testing"
)

type TestStruct struct {
	Field1 string `json:"field_1"`
	Field2 int    `json:"field_2"`
	Field3 bool   `json:"field_3,omitempty"`
	Field4 string
}

func TestGetStructKeys(t *testing.T) {
	expectedKeys := []string{"field_1", "field_2", "field_3"} // Updated expected keys
	result := GetStructKeys(TestStruct{})

	if len(result) != len(expectedKeys) {
		t.Errorf("Test case failed: expected %d keys, got %d keys", len(expectedKeys), len(result))
		return
	}

	for _, key := range expectedKeys {
		if !containsTag(result, key) {
			t.Errorf("Test case failed: expected key %q not found in any struct tag", key)
		}
	}
}

func containsTag(tags []string, key string) bool {
	for _, tag := range tags {
		if tag == key {
			return true
		}
	}
	return false
}

func TestGetFieldName(t *testing.T) {
	tag := "field_2"
	expectedName := "field2" // Updated expected name
	result, err := GetFieldName(tag, TestStruct{})

	if err != nil {
		t.Errorf("Test case failed: unexpected error: %v", err)
	}

	if result != expectedName {
		t.Errorf("Test case failed: expected name=%s, got name=%s", expectedName, result)
	}

	// Test case for non-existent tag
	nonExistentTag := "non_existent"
	_, err = GetFieldName(nonExistentTag, TestStruct{})

	if err == nil || err.Error() != "could not get field name" {
		t.Errorf("Test case failed: expected error message 'could not get field name'")
	}
}
