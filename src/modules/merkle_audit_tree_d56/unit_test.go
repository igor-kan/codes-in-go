package mod_merkle_audit_tree_d56

import "testing"

func TestUnit_Merkle(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	mt := NewMerkleTree([]uint64{10, 20, 30, 40, 50})
	pr := mt.Proof(2)
	if !VerifyProof(30, pr, mt.Root()) || VerifyProof(31, pr, mt.Root()) {
		t.Fatalf("merkle proof verification mismatch")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
