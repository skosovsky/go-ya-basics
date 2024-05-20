package math_test

import (
	"testing"

	math "github.com/skosovsky/go-ya-basics/tests-add"
)

//go:generate go test . -coverprofile=coverage.out
//go:generate go tool cover -html=coverage.out

func TestAddPositive(t *testing.T) {
	t.Parallel()

	sum, err := math.Add(1, 2)
	if err != nil {
		t.Error("unexpected error")
	}
	if sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestAddNegative(t *testing.T) {
	t.Parallel()

	_, err := math.Add(-1, 2)
	if err == nil {
		t.Error("first arg negative - expected error not be nil")
	}
	_, err = math.Add(1, -2)
	if err == nil {
		t.Error("second arg negative - expected error not be nil")
	}
	_, err = math.Add(-1, -2)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
	_, err = math.Add(0, -2)
	if err == nil {
		t.Error("first arg null - expected error not be nil")
	}
	_, err = math.Add(1, 0)
	if err == nil {
		t.Error("second arg null - expected error not be nil")
	}
}
