// utils_test.go
package utils

import (
	"reflect"
	"testing"
)

func TestHexToRGB(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		wantRGB [3]int
		wantErr bool
		errMsg  string
	}{
		{
			name:    "ValidHex",
			hex:     "#FF0000",
			wantRGB: [3]int{255, 0, 0},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "ShortHex",
			hex:     "#F00",
			wantRGB: [3]int{255, 0, 0},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "LowerCaseHex",
			hex:     "#00ff00",
			wantRGB: [3]int{0, 255, 0},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "NoHashHex",
			hex:     "0000FF",
			wantRGB: [3]int{0, 0, 255},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "EmptyHex",
			hex:     "",
			wantRGB: [3]int{},
			wantErr: true,
			errMsg:  "invalid hex color",
		},
		{
			name:    "InvalidCharHex",
			hex:     "#ZZZZZZ",
			wantRGB: [3]int{},
			wantErr: true,
			errMsg:  "invalid hex color",
		},
		{
			name:    "InvalidLengthHex",
			hex:     "#12345",
			wantRGB: [3]int{},
			wantErr: true,
			errMsg:  "invalid hex color",
		},
		{
			name:    "InvalidShortHexLength",
			hex:     "#1234",
			wantRGB: [3]int{},
			wantErr: true,
			errMsg:  "invalid hex color",
		},
		{
			name:    "ValidLowercaseHexWithoutHash",
			hex:     "abcdef",
			wantRGB: [3]int{171, 205, 239},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "HexWithSpecialChars",
			hex:     "#123$%",
			wantRGB: [3]int{},
			wantErr: true,
			errMsg:  "invalid hex color",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRGB, err := HexToRGB(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToRGB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("HexToRGB() error message = %v, wantErrMsg %v", err.Error(), tt.errMsg)
				return
			}
			if !reflect.DeepEqual(gotRGB, tt.wantRGB) {
				t.Errorf("HexToRGB() = %v, want %v", gotRGB, tt.wantRGB)
			}
		})
	}
}
