package mod_booth_min_rotation_go_d61

import "testing"

func TestUnit_BoothRotation(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		s     string
		start int
		want  string
	}{
		{"baaa", 1, "aaab"},
		{"cba", 2, "acb"},
		{"banana", 5, "abanan"},
		{"aaaa", 0, "aaaa"},
		{"abcd", 0, "abcd"},
		{"", 0, ""},
	}
	for _, c := range cases {
		if got := LeastRotation(c.s); got != c.start {
			t.Fatalf("LeastRotation(%q) = %d, want %d", c.s, got, c.start)
		}
		if got := MinimalRotation(c.s); got != c.want {
			t.Fatalf("MinimalRotation(%q) = %q, want %q", c.s, got, c.want)
		}
	}
	info := BoothInfo("banana")
	if info.Start != 5 || info.Text != "abanan" {
		t.Fatalf("unexpected rotation info: %+v", info)
	}
	finder := BoothFinder{}
	if finder.MinimalRotation("cba") != "acb" {
		t.Fatal("finder wrapper mismatch")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
