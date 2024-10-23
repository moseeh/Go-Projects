// utils_test.go
package utils

import (
	"reflect"
	"testing"
)

func TestHSLToRGB(t *testing.T) {
	tests := []struct {
		name    string
		h, s, l int
		want    [3]int
	}{
		{
			name: "Test 1",
			h:    0,
			s:    100,
			l:    50,
			want: [3]int{255, 0, 0}, // Red
		},
		{
			name: "Test 2",
			h:    120,
			s:    100,
			l:    50,
			want: [3]int{0, 255, 0}, // Green
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HSLToRGB(tt.h, tt.s, tt.l)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HSLToRGB() = %v, want %v", got, tt.want)
			}
		})
	}
}
