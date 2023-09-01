package soint

import "testing"

func Test_singleNumberOther2Times(t *testing.T) {
	type args struct {
		nums []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{name: "01", args: args{[]int64{2, 1, 2, 1, 25}}, want: 25},
		{name: "02", args: args{[]int64{-2, 1, -2, 1, -25}}, want: -25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumberOther2Times(tt.args.nums); got != tt.want {
				t.Errorf("singleNumberOther2Times() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleNumberOther3Times(t *testing.T) {
	type args struct {
		nums []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{name: "01", args: args{[]int64{2, 1, 2, 1, 2, 1, 25}}, want: 25},
		{name: "02", args: args{[]int64{-2, 1, -2, 1, -2, 1, -25}}, want: -25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumberOther3Times(tt.args.nums); got != tt.want {
				t.Errorf("singleNumberOther3Times() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mTimesNumberOtherNTimes(t *testing.T) {
	type args struct {
		nums []int64
		n    int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{name: "N: 3, M: 1", args: args{[]int64{2, 1, 2, 1, 2, 1, 25}, 3}, want: 25},
		{name: "N: 3, M: 1", args: args{[]int64{-2, 1, -2, 1, -2, 1, -25}, 3}, want: -25},
		{name: "N: 2, M: 1", args: args{[]int64{2, 1, 2, 1, 25}, 2}, want: 25},
		{name: "N: 2, M: 1", args: args{[]int64{-2, 1, -2, 1, -25}, 2}, want: -25},
		{name: "N: 3, M: 2", args: args{[]int64{2, 1, 2, 1, 2, 1, 25, 25}, 3}, want: 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mTimesNumberOtherNTimes(tt.args.nums, tt.args.n); got != tt.want {
				t.Errorf("mTimesNumberOtherNTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
