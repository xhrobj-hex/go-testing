package even

import "testing"

func TestEven_True(t *testing.T) {
	got := IsEven(2)
	want := true

	if got != want {
		t.Errorf("Even(2) = %v, but want %v", got, want)
	}
}

func TestEven_False(t *testing.T) {
	got := IsEven(1)
	want := false

	if got != want {
		t.Errorf("Even(1) = %v, but want %v", got, want)
	}
}
