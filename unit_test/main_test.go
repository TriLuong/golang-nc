package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(1, 1)
	if result != 2 {
		t.Errorf("1 + 1 expected %d but received %d", 2, result)
	}
}

func TestMul(t *testing.T) {
	result := Mul(1, 1)
	if result != 1 {
		t.Errorf("1 * 1 expected %d but received %d", 1, result)
	}
}
