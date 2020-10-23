package main

import "testing"

func TestMainSuccess(t *testing.T) {

	expected := 0.9999999999999999

	result := loopRaiz()

	if result != expected {
		t.Errorf("Function sum failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("Function sum success")
	}
}
