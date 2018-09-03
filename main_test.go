package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input := 1221

	expected := true

	output := isPalindrome(input)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}

	fmt.Printf("Expected: %v, Output: %v\n", expected, output)
}

func TestCase2(t *testing.T) {
	input := 1234

	expected := false

	output := isPalindrome(input)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}

	fmt.Printf("Expected: %v, Output: %v\n", expected, output)
}
