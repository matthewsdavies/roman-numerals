package main

import "testing"

func Test_convert(t *testing.T) {
	tests := []struct {
		name string
		args int
		want string
	}{
		{
			name: "convert units",
			args: 8,
			want: "VIII",
		},
		{
			name: "convert tens and units",
			args: 34,
			want: "XXXIV",
		},
		{
			name: "convert hundreds, tens and units",
			args: 579,
			want: "DLXXIX",
		},
		{
			name: "convert thousands, hundreds, tens and units",
			args: 2911,
			want: "MMDMXI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
