package mod_singleflight_dedup_group

import (
	"fmt"
	"testing"
)

func TestDistinctKeysExecuteIndependently(t *testing.T) {
	g := NewDedupGroup()
	for i := 0; i < 20; i++ {
		k := fmt.Sprintf("k_%d", i)
		ret := g.Do(k, func() (int64, error) { return int64(i * 10), nil })
		if ret.Val != int64(i*10) || ret.Shared {
			t.Fatalf("unexpected non-shared sequential execution result: %+v", ret)
		}
	}
}
