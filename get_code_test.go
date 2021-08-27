package gauth

import (
	"strings"
	"testing"
)

func TestGetCode(t *testing.T) {
	type args struct {
		secret string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				secret: "MSCSEUJ6ZY7LEJMV",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := GetCode(tt.args.secret)
			t.Logf("code = %v", code)
		})
	}
}

func BenchmarkGetCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		secret := CreateSecret(16)
		code := GetCode(secret)
		if strings.TrimSpace(code) == "" {
			panic("error")
		}
	}
}
