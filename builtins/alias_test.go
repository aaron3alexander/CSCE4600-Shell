package builtins_test

import (
	"errors"
	"testing"

	"github.com/aaron3alexander/CSCE4600-Shell/builtins"
)

func TestAlias(t *testing.T) {

	tests := []struct {
		name    string
		args    []string
		wantErr error
	}{
		{
			name:    "alias listing empty",
			args:    []string{},
			wantErr: errors.New("no aliases"),
		},
		{
			name:    "list non-existent alias",
			args:    []string{"alias1"},
			wantErr: errors.New("alias not found"),
		},
		{
			name:    "add alias",
			args:    []string{"alias1", "pwd"},
			wantErr: nil,
		},
		{
			name:    "list alias",
			args:    []string{},
			wantErr: nil,
		},
		{
			name:    "list specific alias",
			args:    []string{"alias1"},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := builtins.Alias(tt.args...); err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Alias() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}

}
