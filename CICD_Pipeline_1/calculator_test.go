package main

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := Addition(1, 2)
	expected := 3

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(4, 4)
	expected := 16

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestDivision(t *testing.T) {
	result := Divide(30, 3)
	expected := 10

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestDivisionByZero(t *testing.T) {
	result := Divide(20, 0)
	expected := 0

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
