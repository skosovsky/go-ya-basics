package math

import "errors"

var (
	ErrIsZero     = errors.New("arg is zero")
	ErrIsNegative = errors.New("arg is negative")
)

func Add(a, b int) (int, error) { //nolint:varnamelen // example
	if a == 0 || b == 0 {
		return 0, ErrIsZero
	}

	if a < 0 || b < 0 {
		return 0, ErrIsNegative
	}

	return a + b, nil
}
