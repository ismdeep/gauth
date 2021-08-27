package gauth

import (
	"fmt"
	"testing"
)

func TestCreateSecret(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateSecret(16)
			fmt.Println(got)
		})
	}
}
