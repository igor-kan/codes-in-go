package mod_quadtree_spatial_index

import "sync"

const quadCap = 4

type qNode struct {
	box      Rect
	pts      []Point
	children [4]*qNode
	divided  bool
}

type Quadtree struct {
	mu    sync.RWMutex
	root  *qNode
	total int
}

func NewQuadtree(bounds Rect) (*Quadtree, error) {
	if err := validateRect(bounds); err != nil {
		return nil, err
	}
	return &Quadtree{root: &qNode{box: bounds}}, nil
}

func (n *qNode) contains(p Point) bool {
	return p.X >= n.box.MinX && p.X <= n.box.MaxX && p.Y >= n.box.MinY && p.Y <= n.box.MaxY
}

func (r Rect) intersects(o Rect) bool {
	return !(o.MaxX < r.MinX || o.MinX > r.MaxX || o.MaxY < r.MinY || o.MinY > r.MaxY)
}

func (n *qNode) subdivide() {
	mx := (n.box.MinX + n.box.MaxX) * 0.5
	my := (n.box.MinY + n.box.MaxY) * 0.5
	n.children[0] = &qNode{box: Rect{n.box.MinX, my, mx, n.box.MaxY}}
	n.children[1] = &qNode{box: Rect{mx, my, n.box.MaxX, n.box.MaxY}}
	n.children[2] = &qNode{box: Rect{n.box.MinX, n.box.MinY, mx, my}}
	n.children[3] = &qNode{box: Rect{mx, n.box.MinY, n.box.MaxX, my}}
	n.divided = true
}

func (n *qNode) insert(p Point) bool {
	if !n.contains(p) {
		return false
	}
	if !n.divided && len(n.pts) < quadCap {
		n.pts = append(n.pts, p)
		return true
	}
	if !n.divided {
		n.subdivide()
		old := n.pts
		n.pts = nil
		for _, q := range old {
			n.insertIntoChild(q)
		}
	}
	return n.insertIntoChild(p)
}

func (n *qNode) insertIntoChild(p Point) bool {
	for _, ch := range n.children {
		if ch.insert(p) {
			return true
		}
	}
	return false
}

func (q *Quadtree) Insert(p Point) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if !q.root.insert(p) {
		return ErrQuadtreeBounds
	}
	q.total++
	return nil
}

func (q *Quadtree) QueryRange(box Rect) []Point {
	q.mu.RLock()
	defer q.mu.RUnlock()
	var out []Point
	q.root.query(box, &out)
	return out
}

func (n *qNode) query(box Rect, out *[]Point) {
	if !n.box.intersects(box) {
		return
	}
	for _, p := range n.pts {
		if p.X >= box.MinX && p.X <= box.MaxX && p.Y >= box.MinY && p.Y <= box.MaxY {
			*out = append(*out, p)
		}
	}
	if n.divided {
		for _, ch := range n.children {
			ch.query(box, out)
		}
	}
}

func (q *Quadtree) Count() int {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.total
}
