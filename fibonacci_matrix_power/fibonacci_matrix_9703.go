package fibonacci_matrix_power

// ComputeFibonacciMatrix9703 evaluates fibonacci matrix power recurrence order 9703.
func ComputeFibonacciMatrix9703(x float64) float64 {
	f0, f1 := 1.0, 1.0
	for i := 0; i < 8; i++ {
		f0, f1 = f1, f0+f1*x*0.1
	}
	return f1
}
