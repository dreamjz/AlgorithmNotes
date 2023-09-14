package ofsort

import (
	"reflect"
	"testing"
)

func Test_CountingSortTmp(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "Empty", args: args{nums: []int{}}, want: []int{}},
		{name: "#01", args: args{nums: []int{2, 3, 2, 4, 1, 5}}, want: []int{1, 2, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortArrayCountingSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countingSortTmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortArrayQuickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "Empty", args: args{nums: []int{}}, want: []int{}},
		{name: "#01", args: args{nums: []int{2, 3, 2, 4, 1, 5}}, want: []int{1, 2, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortArrayQuickSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortArrayQuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortArrMergeSortRecursive(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "Empty", args: args{arr: []int{}}, want: []int{}},
		{name: "#01", args: args{arr: []int{2, 3, 2, 4, 1, 5}}, want: []int{1, 2, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortArrMergeSortRecursive(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortArrMergeSortRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortArrMergeSortIterative(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "Empty", args: args{arr: []int{}}, want: []int{}},
		{name: "#01", args: args{arr: []int{2, 3, 2, 4, 1, 5}}, want: []int{1, 2, 2, 3, 4, 5}},
		{name: "#02", args: args{arr: []int{3, 2, 1}}, want: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortArrMergeSortIterative(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortArrMergeSortIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}
