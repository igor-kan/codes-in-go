package mod_dinic_blocking_flow_go

import "testing"

func Benchmark_DinicGo(b *testing.B) {
	
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		g := NewDinicGraph(3); g.AddEdge(0, 1, 5); g.AddEdge(1, 2, 5); sink += g.MaxFlow(0, 2)
	}
	_ = sink
}
