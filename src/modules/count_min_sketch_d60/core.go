package mod_count_min_sketch_d60

type CountMinSketch struct {
	d     int
	w     int
	seeds []uint64
	grid  [][]int64
}

func NewCountMinSketch(depth, width int) *CountMinSketch {
	grid := make([][]int64, depth)
	seeds := make([]uint64, depth)
	for r := 0; r < depth; r++ {
		grid[r] = make([]int64, width)
		seeds[r] = uint64(0x9e3779b97f4a7c15) ^ uint64((r+1)*1315423911)
	}
	return &CountMinSketch{d: depth, w: width, seeds: seeds, grid: grid}
}

func (c *CountMinSketch) bucket(row int, key uint64) int {
	x := key ^ c.seeds[row]
	x ^= x >> 33
	x *= 0xff51afd7ed558ccd
	x ^= x >> 33
	return int(x % uint64(c.w))
}

func (c *CountMinSketch) Add(key uint64, count int64) {
	for r := 0; r < c.d; r++ {
		c.grid[r][c.bucket(r, key)] += count
	}
}

func (c *CountMinSketch) Estimate(key uint64) int64 {
	minVal := c.grid[0][c.bucket(0, key)]
	for r := 1; r < c.d; r++ {
		v := c.grid[r][c.bucket(r, key)]
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}
