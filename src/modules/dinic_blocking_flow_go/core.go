package mod_dinic_blocking_flow_go

type DinicGraph struct {
	adj   [][]DinicEdge
	level []int
	ptr   []int
}

func NewDinicGraph(n int) *DinicGraph {
	return &DinicGraph{adj: make([][]DinicEdge, n), level: make([]int, n), ptr: make([]int, n)}
}

func (g *DinicGraph) AddEdge(u, v int, cap int64) {
	g.adj[u] = append(g.adj[u], DinicEdge{To: v, Rev: len(g.adj[v]), Cap: cap})
	g.adj[v] = append(g.adj[v], DinicEdge{To: u, Rev: len(g.adj[u]) - 1, Cap: 0})
}

func (g *DinicGraph) bfs(s, t int) bool {
	for i := range g.level {
		g.level[i] = -1
	}
	g.level[s] = 0
	q := []int{s}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, e := range g.adj[u] {
			if e.Cap > 0 && g.level[e.To] < 0 {
				g.level[e.To] = g.level[u] + 1
				q = append(q, e.To)
			}
		}
	}
	return g.level[t] >= 0
}

func (g *DinicGraph) dfs(u, t int, pushed int64) int64 {
	if pushed == 0 || u == t {
		return pushed
	}
	for ; g.ptr[u] < len(g.adj[u]); g.ptr[u]++ {
		cid := g.ptr[u]
		e := g.adj[u][cid]
		if g.level[e.To] != g.level[u]+1 || e.Cap <= 0 {
			continue
		}
		lim := pushed
		if e.Cap < lim {
			lim = e.Cap
		}
		tr := g.dfs(e.To, t, lim)
		if tr == 0 {
			continue
		}
		g.adj[u][cid].Cap -= tr
		g.adj[e.To][e.Rev].Cap += tr
		return tr
	}
	return 0
}

func (g *DinicGraph) MaxFlow(s, t int) int64 {
	var flow int64
	for g.bfs(s, t) {
		for i := range g.ptr {
			g.ptr[i] = 0
		}
		for {
			pushed := g.dfs(s, t, 1<<60)
			if pushed == 0 {
				break
			}
			flow += pushed
		}
	}
	return flow
}
