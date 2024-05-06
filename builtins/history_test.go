package builtins_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/aaron3alexander/CSCE4600-Shell/builtins"
)

func TestHistory(t *testing.T) {
	// Add commands to history
	builtins.AddHistory("pwd")
	builtins.AddHistory("env")
	builtins.AddHistory("history")

	tests := []struct {
		name    string
		args    []string
		wantErr error
		wantOut string
	}{
		{
			name:    "Print history",
			args:    []string{},
			wantErr: nil,
			wantOut: "1: pwd\n2: env",
		},
		{
			name:    "Clear history",
			args:    []string{"-c"},
			wantErr: nil,
			wantOut: "history cleared",
		},
		{
			name:    "Print empty history",
			args:    nil,
			wantErr: nil,
			wantOut: "history is empty",
		},
		{
			name:    "Invalid arguments",
			args:    []string{"invalid", "args"},
			wantErr: fmt.Errorf("too many arguments"),
			wantOut: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			err := builtins.History(&buf, tt.args...)

			out := strings.TrimSpace(buf.String())
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("History() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("History() error = %v, wantErr %v", err, tt.wantErr)
			}
			if out != tt.wantOut {
				t.Errorf("History() output = %q, want %q", out, tt.wantOut)
			}
		})
	}
}
