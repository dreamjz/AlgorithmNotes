package dsalarr

import "testing"

func Test_findMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "empty array", args: args{nums: []int{}}, want: 0},
		{name: "all 0", args: args{nums: []int{0, 0, 0, 0}}, want: 0},
		{name: "all 1", args: args{nums: []int{1, 1, 1, 1}}, want: 0},
		{name: "01", args: args{nums: []int{0, 1}}, want: 2},
		{name: "02", args: args{nums: []int{0, 1, 0}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("fincMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
