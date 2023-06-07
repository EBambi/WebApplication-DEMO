package main

import "testing"

func TestOrderNumbers(t *testing.T) {
	toOrder := []int{8, 20, 2, 53, 9, 101, 3, 234}
	got := orderNumbers(toOrder)
	want := []int{3, 2, 8, 9, 20, 53, 101, 234}

	if testEq(got, want) == false {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}
