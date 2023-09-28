package golang

import "testing"

func Test_findKthLargest(t *testing.T) {
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
		{name: "#01", args: args{nums: []int{3, 2, 1, 5, 6, 4}, k: 2}, want: 5},
		{name: "#02", args: args{nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4}, want: 4},
		{name: "#03", args: args{nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, k: 2}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
