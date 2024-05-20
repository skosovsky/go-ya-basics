package estimating

func EstimateValue(value int) string {
	switch {
	case value < 10: //nolint:mnd // example
		return "small"
	case value < 100: //nolint:mnd // example
		return "medium"
	default:
		return "big"
	}
}
