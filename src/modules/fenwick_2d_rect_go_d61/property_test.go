package mod_fenwick_2d_rect_go_d61

import "testing"

func TestProp_Fenwick2D(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	grid := [6][5]int64{}
	f := NewFenwick2D(6, 5)
	for i := 0; i < 6; i++ {
		for j := 0; j < 5; j++ {
			v := int64((i*7 + j*3) % 11)
			grid[i][j] = v
			f.Add(i, j, v)
		}
	}
	for x1 := 0; x1 < 6; x1++ {
		for y1 := 0; y1 < 5; y1++ {
			for x2 := x1; x2 < 6; x2++ {
				for y2 := y1; y2 < 5; y2++ {
					var want int64
					for i := x1; i <= x2; i++ {
						for j := y1; j <= y2; j++ {
							want += grid[i][j]
						}
					}
					if got := f.RectSum(RectQuery{X1: x1, Y1: y1, X2: x2, Y2: y2}); got != want {
						t.Fatalf("rect mismatch at %d,%d,%d,%d: got %d want %d", x1, y1, x2, y2, got, want)
					}
				}
			}
		}
	}
}
