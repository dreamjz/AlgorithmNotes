package dsalarr

import "testing"

func Test_numSubarrayLessThanK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "empty array", args: args{nums: []int{}, k: 100}, want: 0},
		{name: "max subarray", args: args{nums: []int{1, 1, 1, 1}, k: 100}, want: 10},
		{name: "01", args: args{nums: []int{10, 5, 2, 6}, k: 100}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayLessThanK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numSubarrayLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}
