package concurrency

import "testing"

func TestFactorialSum(t *testing.T) {
	want := 8
	got := FactorialSum(3)

	if got != want {
		t.Fatal("incorrect value for FactorialSum(8): ", got)
	}
}
