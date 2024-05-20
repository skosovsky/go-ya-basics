package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ErrSomethingWrong = errors.New("something is wrong")

func TestMock_GetData(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		query string
		resp  Response
		err   error
	}{
		{
			name:  "Code 200",
			query: "success",
			resp:  Response{"", 200},
			err:   nil,
		},
		{
			name:  "Code 500",
			query: "error",
			resp:  Response{"", 500},
			err:   ErrSomethingWrong,
		},
		{
			name:  "Code 400",
			query: "empty",
			resp:  Response{"", 400},
			err:   nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var mock Mock
			mock.SetResponse(tt.resp, tt.err)

			resp, err := mock.GetData(tt.query)
			require.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.resp.Text, resp.Text)
		})
	}
}
