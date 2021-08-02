package gauth

import "testing"

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
				secret: "CUDPNHDNMMV6SR34",
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
