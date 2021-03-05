package main

import "testing"

func TestAdd(t *testing.T) {
	var a, b, c int = 3, 8, 11
	result := Add(a, b)
	if c != result {
		t.Errorf("%d != %d", c, result)
	}
}
