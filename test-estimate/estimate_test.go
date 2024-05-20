package estimating_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	estimating "github.com/skosovsky/go-ya-basics/test-estimate"
)

func TestEstimateValue(t *testing.T) {
	t.Parallel()

	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Small",
			args: args{value: 5},
			want: "small",
		},
		{
			name: "Medium",
			args: args{value: 50},
			want: "medium",
		},
		{
			name: "Big",
			args: args{value: 150},
			want: "big",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.EqualValues(t, tt.want, estimating.EstimateValue(tt.args.value))
		})
	}
}
