package mod_radix_patricia_router_go

import "testing"

func TestUnit_RadixRouter(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	rt := NewBytePrefixRouter()
	rt.Insert("/", 1); rt.Insert("/api", 2); rt.Insert("/api/v1/users", 3)
	m := rt.MatchLongest("/api/v1/users/42")
	if !m.Found || m.Handler != 3 {
		t.Fatalf("expected handler 3, got %+v", m)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
