package dsalarr

import "testing"

func Test_subarraySumBruteForce(t *testing.T) {
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
		{name: "empty array", args: args{nums: []int{}, k: 2}, want: 0},
		{name: "not found", args: args{nums: []int{1, 1, 1, 1}, k: 100}, want: 0},
		{name: "all elem is k", args: args{nums: []int{1, 1, 1, 1}, k: 1}, want: 4},
		{name: "01", args: args{nums: []int{1, 1, 1, 1}, k: 2}, want: 3},
		{name: "02", args: args{nums: []int{-1, -1, 1}, k: 0}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySumBruteForce(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySumBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subarraySumBruteForceOptimized(t *testing.T) {
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
		{name: "empty array", args: args{nums: []int{}, k: 2}, want: 0},
		{name: "not found", args: args{nums: []int{1, 1, 1, 1}, k: 100}, want: 0},
		{name: "all elem is k", args: args{nums: []int{1, 1, 1, 1}, k: 1}, want: 4},
		{name: "01", args: args{nums: []int{1, 1, 1, 1}, k: 2}, want: 3},
		{name: "02", args: args{nums: []int{-1, -1, 1}, k: 0}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySumBruteForceOptimized(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySumBruteForceOptimized() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subarraySumHashTable(t *testing.T) {
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
		{name: "empty array", args: args{nums: []int{}, k: 2}, want: 0},
		{name: "not found", args: args{nums: []int{1, 1, 1, 1}, k: 100}, want: 0},
		{name: "all elem is k", args: args{nums: []int{1, 1, 1, 1}, k: 1}, want: 4},
		{name: "01", args: args{nums: []int{1, 1, 1, 1}, k: 2}, want: 3},
		{name: "02", args: args{nums: []int{-1, -1, 1}, k: 0}, want: 1},
		{name: "03", args: args{nums: []int{1, 2, 3}, k: 3}, want: 2},
		{name: "04", args: args{nums: []int{1, 1, 1}, k: 2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySumHashTable(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySumHashTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
