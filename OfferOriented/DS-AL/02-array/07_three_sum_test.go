package dsalarr

import (
	"reflect"
	"testing"
)

func Test_threeSumHashTable(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{name: "empty array", args: args{nums: []int{}}, want: [][]int{}},
		{name: "not found", args: args{nums: []int{4, -2, 4, 2, 5}}, want: [][]int{}},
		{name: "all 0", args: args{nums: []int{0, 0, 0, 0, 0}}, want: [][]int{{0, 0, 0}}},
		{name: "01", args: args{nums: []int{-1, 0, 1, 2, -1, 4}}, want: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSumHashTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
