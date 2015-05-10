package main

import "testing"

//FizzBuzzResult is used for testing
type FizzBuzzResult struct {
	input    int
	expected string
}

var fizzBuzzTests = []FizzBuzzResult{
	{1, "1"},
	{3, "Fizz"},
	{5, "Buzz"},
	{15, "FizzBuzz"},
	{7, "7"},
}

func TestFizzBuzz(t *testing.T) {
	for _, fbr := range fizzBuzzTests {
		restult := FizzBuzz(fbr.input).String()
		if restult != fbr.expected {
			t.Error("Expected "+fbr.expected+" got ", restult)
		}
	}
}
