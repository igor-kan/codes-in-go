package mod_douglas_peucker_go_d61

import "math"

func PerpendicularDistance(p, a, b Point) float64 {
	dx := b.X - a.X
	dy := b.Y - a.Y
	if dx == 0 && dy == 0 {
		return math.Hypot(p.X-a.X, p.Y-a.Y)
	}
	num := math.Abs(dy*p.X - dx*p.Y + b.X*a.Y - b.Y*a.X)
	return num / math.Hypot(dx, dy)
}

func SimplifyPolyline(points []Point, epsilon float64) []Point {
	if len(points) <= 2 {
		out := make([]Point, len(points))
		copy(out, points)
		return out
	}
	keep := make([]bool, len(points))
	keep[0] = true
	keep[len(points)-1] = true
	stack := [][2]int{{0, len(points) - 1}}
	for len(stack) > 0 {
		seg := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		a, b := seg[0], seg[1]
		maxDist := -1.0
		idx := -1
		for i := a + 1; i < b; i++ {
			d := PerpendicularDistance(points[i], points[a], points[b])
			if d > maxDist {
				maxDist = d
				idx = i
			}
		}
		if idx != -1 && maxDist > epsilon {
			keep[idx] = true
			stack = append(stack, [2]int{a, idx}, [2]int{idx, b})
		}
	}
	out := make([]Point, 0, len(points))
	for i, k := range keep {
		if k {
			out = append(out, points[i])
		}
	}
	return out
}

type RdpSimplifier struct{}

func (r RdpSimplifier) Simplify(points []Point, epsilon float64) []Point {
	return SimplifyPolyline(points, epsilon)
}

var _ PolylineSimplifier = RdpSimplifier{}
