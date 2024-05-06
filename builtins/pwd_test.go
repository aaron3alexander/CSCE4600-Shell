package builtins_test

import (
	"testing"

	"github.com/aaron3alexander/CSCE4600-Shell/builtins"
)

func Test_Pwd(t *testing.T) {
	tests := []struct {
		name    string
		wantErr error
	}{
		{
			name:    "pwd",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := builtins.PrintWorkingDirectory(); tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("PrintWorkingDirectory() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
