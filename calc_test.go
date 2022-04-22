package main


// Import the testing Package 
import (
	"testing"
)


// Import the testing Package 
func TestCalculate(t *testing.T) {

	// Function named and expected output
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

