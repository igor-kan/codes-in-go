package mod_interval_coalesce_go

func validateSpan(s Span) error {
	if s.Start > s.End {
		return ErrIntervalInvalid
	}
	return nil
}
