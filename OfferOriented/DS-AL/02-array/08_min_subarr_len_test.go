package dsalarr

import "testing"

func Test_minSubArrayLen(t *testing.T) {
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
		{name: "empty array", args: args{nums: []int{}, k: 7}, want: 0},
		{name: "max subarray", args: args{nums: []int{1, 1, 1, 1}, k: 4}, want: 4},
		{name: "01", args: args{nums: []int{5, 1, 4, 3}, k: 7}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
