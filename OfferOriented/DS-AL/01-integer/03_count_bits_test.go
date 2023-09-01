package soint

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func Test_countBitsOkOfOneNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "01", args: args{0}, want: 0},
		{name: "02", args: args{1}, want: 1},
		{name: "03", args: args{2}, want: 1},
		{name: "03", args: args{3}, want: 2},
		{name: "03", args: args{4}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBitsOkOfOneNum(tt.args.num); got != tt.want {
				t.Errorf("countBitsOnk(%d) = %v, want %v", tt.args.num, got, tt.want)
			}
		})
	}
}

func Test_countBitsOnk(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "01", args: args{4}, want: []int{0, 1, 1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBitsOnk(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBitsOnk(%d) = %v, want %v", tt.args.num, got, tt.want)
			}
		})
	}
}

func Test_countBitsOn(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "01", args: args{4}, want: []int{0, 1, 1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBitsOn(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBitsOn(%d) = %v, want %v", tt.args.num, got, tt.want)
			}
		})
	}
}

func Test_countBitsOnWithBitOperation(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "01", args: args{4}, want: []int{0, 1, 1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBitsOnWithBitOperation(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBitsOnWithBitOperation(%d) = %v, want %v", tt.args.num, got, tt.want)
			}
		})
	}
}

type generator struct {
	src *rand.Rand
}

func (g *generator) nextInt(max int) int {
	return g.src.Intn(max)
}

func testSet(seed, count int) []int {
	gen := &generator{
		src: rand.New(rand.NewSource(int64(seed))),
	}

	set := make([]int, count)
	for i := range set {
		set[i] = gen.nextInt(count)
	}
	return set
}

func BenchmarkCountBitsOnk(b *testing.B) {
	for n := 10; n <= 100000; n *= 10 {
		set := testSet(300, n)
		b.ResetTimer()
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, v := range set {
					countBitsOnk(v)
				}
			}
		})

	}
}

func BenchmarkCountBitsOn(b *testing.B) {
	for n := 10; n <= 100000; n *= 10 {
		set := testSet(300, n)
		b.ResetTimer()
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, v := range set {
					countBitsOn(v)
				}
			}
		})

	}
}
func BenchmarkCountBitsOnWithBitOperation(b *testing.B) {
	for n := 10; n <= 100000; n *= 10 {
		set := testSet(300, n)
		b.ResetTimer()
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, v := range set {
					countBitsOnWithBitOperation(v)
				}
			}
		})

	}
}
