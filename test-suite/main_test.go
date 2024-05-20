package main_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	StartNumFive int
}

func (s *TestSuite) SetupTest() {
	s.StartNumFive = 5
}

func (s *TestSuite) Test() {
	s.Equal(5, s.StartNumFive)
}

func TestTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(TestSuite))
}
