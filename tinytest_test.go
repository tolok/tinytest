package tinytest

import (
	"testing"
)

func TestMapValues(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test MapValues",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapValues()
			for k, v := range got {
				if v != k*2+1 {
					t.Errorf("unexpected value")
				}
			}
		})
	}
}
