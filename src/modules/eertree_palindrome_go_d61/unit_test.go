package mod_eertree_palindrome_go_d61

import "testing"

func TestUnit_Eertree(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		s    string
		want int
	}{
		{"abba", 4},
		{"aaa", 3},
		{"abacaba", 7},
		{"abcd", 4},
		{"", 0},
	}
	for _, c := range cases {
		if got := DistinctPalindromes(c.s); got != c.want {
			t.Fatalf("DistinctPalindromes(%q) = %d, want %d", c.s, got, c.want)
		}
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
