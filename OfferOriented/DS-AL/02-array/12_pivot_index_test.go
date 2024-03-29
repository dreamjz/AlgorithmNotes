package dsalarr

import "testing"

func Test_pivotIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "empty array", args: args{nums: []int{}}, want: -1},
		{name: "01", args: args{nums: []int{1, 7, 3, 6, 5, 6}}, want: 3},
		{name: "02", args: args{nums: []int{1, 2, 3}}, want: -1},
		{name: "02", args: args{nums: []int{2, 1, -1}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
