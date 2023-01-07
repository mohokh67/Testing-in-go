package calculator

import "testing"

func TestSum(t *testing.T) {

	got := Sum(3.0, 6.0)
	expected := 9.0

	// if 2 != 3 {
	// 	t.Error("silly assertion, we know 2 and 3 are not equal")
	// }

	if got != expected {
		t.Errorf("Did not get expected result. Wanted %v, got: %v\n", expected, got)
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	got := Subtract(8.0, 3.0)
	expected := 5.0

	if got != expected {
		t.Errorf("Did not get expected result. Wanted %v, got: %v\n", expected, got)
	}
}

func TestSubtractTableDriven(t *testing.T) {
	t.Log("this is table driven test")
	t.Parallel()
	scenarios := []struct {
		num1   float64
		num2   float64
		expect float64
	}{
		{num1: 20, num2: 7, expect: 13},
		{num1: 5, num2: 2, expect: 3},
		{num1: 12, num2: 20, expect: -8},
	}

	for _, s := range scenarios {
		got := Subtract(s.num1, s.num2)

		if got != s.expect {
			t.Errorf("Did not get expected result for input %v and %v. Wanted %v, got: %v\n", s.num1, s.num2, s.expect, got)
		}
	}

}

func TestFailures(t *testing.T) {
	t.Error("Failed but move on")
	t.Fatal("Fatal error and the test will stop here")
	t.Error("This line never be printed")
}
