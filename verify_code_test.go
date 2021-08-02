package gauth

import (
	"math/rand"
	"testing"
	"time"
)

func TestVerifyCode(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	secret := CreateSecret()
	code := GetCode(secret)

	type args struct {
		secret string
		code   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				secret: secret,
				code:   code,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyCode(tt.args.secret, tt.args.code); got != tt.want {
				t.Errorf("VerifyCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
