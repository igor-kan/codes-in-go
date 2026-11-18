package mod_jump_consistent_hash_go

import "testing"

func TestUnit_JumpHash(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	for k := uint64(0); k < 100; k++ {
		b10 := JumpConsistentHash(k, 10)
		b11 := JumpConsistentHash(k, 11)
		if b10 < 0 || b10 >= 10 || (b11 != b10 && b11 != 10) {
			t.Fatalf("monotonicity broken: b10=%d b11=%d", b10, b11)
		}
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
