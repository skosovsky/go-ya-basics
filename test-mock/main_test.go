package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type DBMock struct {
	emails map[string]bool
}

func (d DBMock) UserExists(email string) bool {
	return d.emails[email]
}

func (d DBMock) addUser(email string) {
	d.emails[email] = true
}

func TestNewUser(t *testing.T) {
	t.Parallel()

	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		preset  bool
		wantErr bool
	}{
		{
			name: "want success",
			args: args{
				email: "gregorysmith@myexampledomain.com",
			},
			preset:  false,
			wantErr: false,
		},
		{
			name: "want error",
			args: args{
				email: "johndoe@myexampledomain.com",
			},
			preset:  true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dbFake := &DBMock{emails: make(map[string]bool)}
			if tt.preset {
				dbFake.addUser(tt.args.email)
			}

			err := NewUser(dbFake, tt.args.email)

			if !tt.wantErr {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, ErrUserExists)
			}
		})
	}
}
