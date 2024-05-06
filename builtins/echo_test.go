package builtins_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/aaron3alexander/CSCE4600-Shell/builtins"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr error
		wantOut string
	}{
		{
			name:    "Echo single word",
			args:    []string{"hello"},
			wantErr: nil,
			wantOut: "hello\n",
		},
		{
			name:    "Echo multiple words",
			args:    []string{"hello", "world"},
			wantErr: nil,
			wantOut: "hello world\n",
		},
		{
			name:    "Empty echo",
			args:    []string{},
			wantErr: errors.New("no arguments given"),
			wantOut: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			err := builtins.Echo(&buf, tt.args...)

			out := buf.String()
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("Echo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("Echo() error = %v, wantErr %v", err, tt.wantErr)
			}
			if out != tt.wantOut {
				t.Errorf("Echo() output = %q, want %q", out, tt.wantOut)
			}
		})
	}
}
