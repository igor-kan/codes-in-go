package mod_fenwick_2d_rect_go_d61

type Fenwick2D struct {
	rows int
	cols int
	tree [][]int64
}

func NewFenwick2D(rows, cols int) *Fenwick2D {
	tree := make([][]int64, rows+1)
	for i := range tree {
		tree[i] = make([]int64, cols+1)
	}
	return &Fenwick2D{rows: rows, cols: cols, tree: tree}
}

func (f *Fenwick2D) Add(x, y int, delta int64) {
	for i := x + 1; i <= f.rows; i += i & -i {
		for j := y + 1; j <= f.cols; j += j & -j {
			f.tree[i][j] += delta
		}
	}
}

func (f *Fenwick2D) prefix(x, y int) int64 {
	var sum int64
	for i := x + 1; i > 0; i -= i & -i {
		for j := y + 1; j > 0; j -= j & -j {
			sum += f.tree[i][j]
		}
	}
	return sum
}

func (f *Fenwick2D) RectSum(q RectQuery) int64 {
	if q.X1 > q.X2 || q.Y1 > q.Y2 {
		return 0
	}
	s := f.prefix(q.X2, q.Y2)
	if q.X1 > 0 {
		s -= f.prefix(q.X1-1, q.Y2)
	}
	if q.Y1 > 0 {
		s -= f.prefix(q.X2, q.Y1-1)
	}
	if q.X1 > 0 && q.Y1 > 0 {
		s += f.prefix(q.X1-1, q.Y1-1)
	}
	return s
}

var _ GridRangeSummer = (*Fenwick2D)(nil)
