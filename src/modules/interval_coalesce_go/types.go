package mod_interval_coalesce_go

type Span struct {
	Start int64
	End   int64
}

func (s Span) Width() int64 {
	return s.End - s.Start + 1
}
