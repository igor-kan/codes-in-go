package mod_count_min_heavy_go

func validateSketchDims(depth, width int) error {
	if depth <= 0 || width <= 0 {
		return ErrCountMinInvalid
	}
	return nil
}
