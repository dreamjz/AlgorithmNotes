package soint

import (
	"testing"
)

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		args args
		want string
	}{
		// TODO: Add test cases.
		{args: args{"00", "00"}, want: "00"},
		{args: args{"01", "01"}, want: "10"},
		{args: args{"11", "11"}, want: "110"},
		{args: args{"10101010", "11001100"}, want: "101110110"},
	}
	for _, tt := range tests {
		if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
			t.Errorf("addBinary(%s, %s) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
		}
	}
}
