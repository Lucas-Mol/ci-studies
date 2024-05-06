package main

import "testing"

func TestSoma(t *testing.T) {
	total := Sum(15, 15)

	if total != 30 {
		t.Errorf("The sum result is incorrect: Result %d, Must be %d", total, 30)
	}
}
