package mod_bloom_double_hash_d56

import "testing"

func TestUnit_Bloom(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	bf := NewBloomFilter(256, 4)
	bf.Add(1001)
	bf.Add(2002)
	if !bf.MightContain(1001) || !bf.MightContain(2002) {
		t.Fatalf("false negative detected")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
