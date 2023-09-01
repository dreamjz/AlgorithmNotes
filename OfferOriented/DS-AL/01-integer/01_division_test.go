package soint

import (
	"math"
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		dividend int32
		divisor  int32
	}
	tests := []struct {
		args args
		want int32
	}{
		// TODO: Add test cases.
		{args: args{dividend: 12, divisor: 6}, want: 2},
		{args: args{dividend: 15, divisor: 2}, want: 7},
		{args: args{dividend: -15, divisor: 2}, want: -7},
		{args: args{dividend: 15, divisor: -2}, want: -7},
		{args: args{dividend: -15, divisor: -2}, want: 7},
		{args: args{dividend: math.MaxInt32, divisor: 1}, want: math.MaxInt32},
		{args: args{dividend: math.MinInt32, divisor: -1}, want: math.MaxInt32},
	}
	for _, tt := range tests {
		if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
			t.Errorf("divide() = %v, want %v", got, tt.want)
		}
	}
}
