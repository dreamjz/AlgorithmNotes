package dsalarr

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func Test_twoSumBruteForce(t *testing.T) {
	type args struct {
		ints []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "01", args: args{ints: []int{1, 2, 4, 6, 10}, n: 8}, want: []int{1, 3}},
		{name: "02", args: args{ints: []int{1, 2, 4, 6, 10}, n: 20}, want: []int{}},
		{name: "empty array", args: args{ints: []int{}, n: 20}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumBruteForce(tt.args.ints, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSumBinarySearch(t *testing.T) {
	type args struct {
		ints []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "01", args: args{ints: []int{1, 2, 4, 6, 10}, n: 8}, want: []int{1, 3}},
		{name: "02", args: args{ints: []int{1, 2, 4, 6, 10}, n: 20}, want: []int{}},
		{name: "empty array", args: args{ints: []int{}, n: 20}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumBinarySearch(tt.args.ints, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumBinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearch(t *testing.T) {
	type args struct {
		ints []int
		t    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "01", args: args{ints: []int{-1, 0, 3, 5, 9, 12}, t: 2}, want: -1},
		{name: "02", args: args{ints: []int{1, 2, 4, 6, 10}, t: 6}, want: 3},
		{name: "empty array", args: args{ints: []int{}, t: 20}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.ints, tt.args.t); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSumHashTable(t *testing.T) {
	type args struct {
		ints []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "01", args: args{ints: []int{1, 2, 4, 6, 10}, n: 8}, want: []int{1, 3}},
		{name: "02", args: args{ints: []int{1, 2, 4, 6, 10}, n: 20}, want: []int{}},
		{name: "empty array", args: args{ints: []int{}, n: 20}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumHashTable(tt.args.ints, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumHashTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSumTwoPointer(t *testing.T) {
	type args struct {
		ints []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "01", args: args{ints: []int{1, 2, 4, 6, 10}, n: 8}, want: []int{1, 3}},
		{name: "02", args: args{ints: []int{1, 2, 4, 6, 10}, n: 20}, want: []int{}},
		{name: "empty array", args: args{ints: []int{}, n: 20}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumTwoPointer(tt.args.ints, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumTwoPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

type generator06 struct {
	src *rand.Rand
}

func (g *generator06) nextInt(max, min int) int {
	return g.src.Intn(max-min+1) + min
}

func testSet(size int, max, min int) []int {
	gen := &generator06{
		src: rand.New(rand.NewSource(time.Now().Unix())),
	}

	ints := make([]int, size)
	for i := range ints {
		ints[i] = gen.nextInt(max, min)
	}

	return ints
}

// ascending order
func sort(ints []int) {
	quickSort(ints, 0, len(ints)-1)
}

func quickSort(ints []int, left, right int) {
	if left < right {
		less, great := partition(ints, left, right)
		quickSort(ints, left, less-1)
		quickSort(ints, great+1, right)
	}
}

func partition(ints []int, left, right int) (int, int) {
	// random pivot
	pIdx := rand.Intn(right-left+1) + left
	p := ints[pIdx]
	swap(ints, pIdx, right)

	// 3-partiton: Less, Equal, Great
	less := left
	great := right
	idx := left
	for idx <= great {
		ele := ints[idx]
		if ele == p {
			idx++
		}
		if ele < p {
			swap(ints, idx, less)
			idx++
			less++
		}
		if ele > p {
			swap(ints, idx, great)
			great--
		}
	}

	return less, great
}

func swap(ints []int, i, j int) {
	ints[i], ints[j] = ints[j], ints[i]
}

func BenchmarkTwoSum(b *testing.B) {
	tests := []struct {
		name string
		f    func(nums []int, n int) []int
	}{
		{name: "BruteForce", f: twoSumBruteForce},
		{name: "BinarySearch", f: twoSumBinarySearch},
		{name: "HashTable", f: twoSumHashTable},
		{name: "TwoPointer", f: twoSumTwoPointer},
	}

	for k := 1000; k <= 100000; k *= 10 {
		set := testSet(k, k, -k)
		sort(set)
		target := rand.Intn(2*k+1) - k
		b.ResetTimer()
		for _, tt := range tests {
			b.Run(fmt.Sprintf("%s-%.E", tt.name, float64(k)), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tt.f(set, target)
				}
			})
		}
	}
}
