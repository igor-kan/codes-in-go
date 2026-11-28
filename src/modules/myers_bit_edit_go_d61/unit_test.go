package mod_myers_bit_edit_go_d61

import "testing"

func TestUnit_MyersEdit(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	if got := MyersDistance("kitten", "sitting"); got != 3 {
		t.Fatalf("kitten/sitting = %d, want 3", got)
	}
	if got := MyersDistance("", "abc"); got != 3 {
		t.Fatalf("empty pattern = %d, want 3", got)
	}
	if got := MyersDistance("abc", ""); got != 3 {
		t.Fatalf("empty text = %d, want 3", got)
	}
	if got := MyersDistance("abc", "abc"); got != 0 {
		t.Fatalf("identical strings = %d, want 0", got)
	}
	if got := MyersDistance("ab", "ba"); got != 2 {
		t.Fatalf("swap = %d, want 2", got)
	}
	longPat := "0123456789012345678901234567890123456789012345678901234567890123456789"
	if got := MyersDistance(longPat, longPat); got != 0 {
		t.Fatalf("long fallback mismatch = %d, want 0", got)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
