package mod_radix_patricia_router_go

type trieNode struct {
	next    map[byte]*trieNode
	set     bool
	prefix  string
	handler int
}

type BytePrefixRouter struct {
	root *trieNode
}

func NewBytePrefixRouter() *BytePrefixRouter {
	return &BytePrefixRouter{root: &trieNode{next: make(map[byte]*trieNode)}}
}

func (r *BytePrefixRouter) Insert(prefix string, handler int) {
	cur := r.root
	for i := 0; i < len(prefix); i++ {
		b := prefix[i]
		if cur.next[b] == nil {
			cur.next[b] = &trieNode{next: make(map[byte]*trieNode)}
		}
		cur = cur.next[b]
	}
	cur.set = true
	cur.prefix = prefix
	cur.handler = handler
}

func (r *BytePrefixRouter) MatchLongest(path string) RouteMatch {
	cur := r.root
	var best RouteMatch
	if cur.set {
		best = RouteMatch{Prefix: cur.prefix, Handler: cur.handler, Found: true}
	}
	for i := 0; i < len(path); i++ {
		cur = cur.next[path[i]]
		if cur == nil {
			break
		}
		if cur.set {
			best = RouteMatch{Prefix: cur.prefix, Handler: cur.handler, Found: true}
		}
	}
	return best
}
