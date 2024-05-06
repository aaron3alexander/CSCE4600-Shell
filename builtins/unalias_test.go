package builtins_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/aaron3alexander/CSCE4600-Shell/builtins"
)

func TestUnAlias(t *testing.T) {
	if err := builtins.Alias("alias1", "pwd"); err != nil {
		t.Fatalf("Error setting up test: %v", err)
	}
	if err := builtins.Alias("alias2", "echo hello"); err != nil {
		t.Fatalf("Error setting up test: %v", err)
	}
	tests := []struct {
		name    string
		args    []string
		wantErr error
		wantOut string
	}{
		{
			name:    "Remove existing alias",
			args:    []string{"alias2"},
			wantErr: nil,
			wantOut: "Removed alias: alias2 \n",
		},
		{
			name:    "Remove non-existent alias",
			args:    []string{"random"},
			wantErr: nil,
			wantOut: "Alias 'random' not found\n",
		},
		{
			name:    "Remove multiple aliases",
			args:    []string{"alias1", "alias2"},
			wantErr: nil,
			wantOut: "Removed alias: alias1 \nAlias 'alias2' not found\n",
		},
		{
			name:    "No alias name",
			args:    []string{},
			wantErr: fmt.Errorf("missing alias name"),
			wantOut: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			err := builtins.UnAlias(&buf, tt.args...)

			out := buf.String()
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("UnAlias() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("UnAlias() error = %v, wantErr %v", err, tt.wantErr)
			}
			if out != tt.wantOut {
				t.Errorf("UnAlias() output = %q, want %q", out, tt.wantOut)
			}
		})
	}
}
