package stringee

import "testing"

func TestTrimString(t *testing.T) {
	testCases := []struct {
		value  string
		trim   int
		expect string
	}{
		{value: "abc", trim: 1, expect: "ab"},
		{value: "aa", trim: 0, expect: "aa"},
		{value: "hello", trim: 2, expect: "hel"},
	}

	for _, s := range testCases {
		got := TrimString(s.value, s.trim)

		if got != s.expect {
			t.Errorf("Did not get expected result when trim %v character(s) from %v. Wanted %v, got: %v\n", s.trim, s.value, s.expect, got)
		}
	}
}

func FuzzTrimString(f *testing.F) {

	testCases := []struct {
		value string
		trim  int
	}{
		{value: "abc", trim: 1},
		{value: "aa", trim: 0},
		{value: "hello", trim: 2},
	}

	for _, s := range testCases {
		f.Add(s.value, s.trim)
	}
	f.Fuzz(func(t *testing.T, value string, num int) {
		TrimString(value, num)
	})

}
