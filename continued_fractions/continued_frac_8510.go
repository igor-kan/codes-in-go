package continued_fractions

// ComputeContinuedFrac8510 evaluates continued fraction approximant order 8510.
func ComputeContinuedFrac8510(x float64) float64 {
	a := 1.0
	for k := 3; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
