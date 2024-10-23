package utils

import "testing"

func TestCheckArgs(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "One argument, not a flag",
			args: args{args: []string{"hello"}},
			want: true,
		},
		{
			name: "One argument, is a flag",
			args: args{args: []string{"--help"}},
			want: false,
		},
		{
			name: "Two arguments, flag and string",
			args: args{args: []string{"--reverse", "hello"}},
			want: true,
		},
		{
			name: "Two arguments, string and banner",
			args: args{args: []string{"hello", "standard"}},
			want: true,
		},
		{
			name: "Two arguments, invalid banner",
			args: args{args: []string{"hello", "invalidBanner"}},
			want: false,
		},
		{
			name: "Three arguments, flag, string, valid banner",
			args: args{args: []string{"--reverse", "hello", "standard"}},
			want: true,
		},
		{
			name: "Three arguments, flag, string, invalid banner",
			args: args{args: []string{"--reverse", "hello", "invalidBanner"}},
			want: false,
		},
		{
			name: "Three arguments, color flag, string, string",
			args: args{args: []string{"--color=red", "hello", "world"}},
			want: true,
		},
		{
			name: "Four arguments, color flag, string, string, valid banner",
			args: args{args: []string{"--color=red", "hello", "world", "standard"}},
			want: true,
		},
		{
			name: "Four arguments, color flag, string, string, invalid banner",
			args: args{args: []string{"--color=red", "hello", "world", "invalidBanner"}},
			want: false,
		},
		{
			name: "Five arguments, invalid case",
			args: args{args: []string{"--color=red", "hello", "world", "standard", "extra"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckArgs(tt.args.args); got != tt.want {
				t.Errorf("CheckArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
