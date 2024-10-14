package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected float64
	}{
		{1, 2, 3},
		{-1, -1, -2},
		{1.5, 2.5, 4},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b, expected float64
	}{
		{3, 2, 1},
		{-1, -1, 0},
		{5.5, 2.5, 3},
	}

	for _, tt := range tests {
		result := Subtract(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Subtract(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, expected float64
	}{
		{3, 2, 6},
		{-1, -1, 1},
		{1.5, 2.5, 3.75},
	}

	for _, tt := range tests {
		result := Multiply(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Multiply(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b, expected float64
		expectError    bool
	}{
		{6, 2, 3, false},
		{-6, -2, 3, false},
		{6, 0, 0, true},
		{5.5, 2.5, 2.2, false},
	}

	for _, tt := range tests {
		result, err := Divide(tt.a, tt.b)
		if (err != nil) != tt.expectError {
			t.Errorf("Divide(%f, %f) returned error %v; expectError %v", tt.a, tt.b, err, tt.expectError)
		}
		if !tt.expectError && result != tt.expected {
			t.Errorf("Divide(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
		}
	}
}
