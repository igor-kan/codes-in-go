package mod_count_min_sketch_d57

type CountMin struct {
	depth int
	width uint64
	table [][]uint64
}

func NewCountMin(width int) *CountMin {
	tbl := make([][]uint64, 4)
	for i := 0; i < 4; i++ {
		tbl[i] = make([]uint64, width)
	}
	return &CountMin{depth: 4, width: uint64(width), table: tbl}
}

func (c *CountMin) hashRow(item uint64, row int) uint64 {
	salt := uint64(row+1) * 0x9E3779B97F4A7C15
	x := (item ^ salt) * 0xBF58476D1CE4E5B9
	return (x ^ (x >> 29)) % c.width
}

func (c *CountMin) Add(item, delta uint64) {
	for r := 0; r < c.depth; r++ {
		col := c.hashRow(item, r)
		c.table[r][col] += delta
	}
}

func (c *CountMin) Query(item uint64) uint64 {
	best := c.table[0][c.hashRow(item, 0)]
	for r := 1; r < c.depth; r++ {
		v := c.table[r][c.hashRow(item, r)]
		if v < best {
			best = v
		}
	}
	return best
}
