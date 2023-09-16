package ofdp

import "testing"

func Test_lenLongestFibSubseq(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "#01", args: args{arr: []int{1, 3, 7, 11, 12, 14, 18}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lenLongestFibSubseq(tt.args.arr); got != tt.want {
				t.Errorf("lenLongestFibSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
