package magical_test

import (
	"testing"

	"github.com/mohokh67/go-testing/magical"
)

func TestSomethingElse(t *testing.T) {
	got := magical.DoSomethingElse()
	expected := false

	if got != expected {
		t.Errorf("Did not get expected result. Wanted %v, got: %v\n", expected, got)
	}
}

func TestAdd(t *testing.T) {
	t.Skip("we will ignore this test for now")
	got := 3 + 5
	expected := 10

	if got != expected {
		t.Errorf("Did not get expected result. Wanted %v, got: %v\n", expected, got)
	}
}
