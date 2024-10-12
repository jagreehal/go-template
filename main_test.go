package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

// Test a successful case with the input "World"
func TestHelloWorld(t *testing.T) {
	result, err := Hello("World")
	assert.NoError(t, err, "Expected no error for valid input")
	assert.Equal(t, "Hello, World!", result)
}

// Test another successful case with the input "Go"
func TestHelloGo(t *testing.T) {
	expected := "Hello, Go!"
	actual, err := Hello("Go")
	assert.NoError(t, err, "Expected no error for valid input")
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("Hello() mismatch (-want +got):\n%s", diff)
	}
}

// Test with an empty name to ensure proper error handling
func TestHelloEmptyName(t *testing.T) {
	_, err := Hello("")
	assert.Error(t, err, "Expected an error for empty name")
	assert.Contains(t, err.Error(), "name cannot be empty", "Error message should indicate the name is missing")
}
