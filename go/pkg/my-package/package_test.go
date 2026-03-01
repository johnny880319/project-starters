package mypackage

import "testing"

func TestMyFunction1(t *testing.T) {
	t.Parallel()
	expected := "Hello, World!"
	result := MyFunction1()
	if result != expected {
		t.Errorf("MyFunction1() = %v; want %v", result, expected)
	}
}

func TestMyFunction2(t *testing.T) {
	t.Parallel()
	expected := 42
	result := MyFunction2()
	if result != expected {
		t.Errorf("MyFunction2() = %v; want %v", result, expected)
	}
}
