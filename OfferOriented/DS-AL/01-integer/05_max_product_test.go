package soint

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func Test_maxProductBruteForce(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "empty string array", args: args{strs: []string{}}, want: 0},
		{name: "contains empty element", args: args{strs: []string{"ab", "", "ef", "", "qwe", "asd"}}, want: 9},
		{name: "01", args: args{strs: []string{"ab", "bd", "ef", "pq", "qwe", "asd"}}, want: 9},
		{name: "02", args: args{strs: []string{"asb", "bsd", "dse", "esf", "fsg", "gsh"}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductBruteForce(tt.args.strs); got != tt.want {
				t.Errorf("maxProductBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProductHashTable(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "empty string array", args: args{strs: []string{}}, want: 0},
		{name: "contains empty element", args: args{strs: []string{"ab", "", "ef", "", "qwe", "asd"}}, want: 9},
		{name: "01", args: args{strs: []string{"ab", "bd", "ef", "pq", "qwe", "asd"}}, want: 9},
		{name: "02", args: args{strs: []string{"asb", "bsd", "dse", "esf", "fsg", "gsh"}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductHashTable(tt.args.strs); got != tt.want {
				t.Errorf("maxProductHashTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProductBinary(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "empty string array", args: args{strs: []string{}}, want: 0},
		{name: "contains empty element", args: args{strs: []string{"ab", "", "ef", "", "qwe", "asd"}}, want: 9},
		{name: "01", args: args{strs: []string{"ab", "bd", "ef", "pq", "qwe", "asd"}}, want: 9},
		{name: "02", args: args{strs: []string{"asb", "bsd", "dse", "esf", "fsg", "gsh"}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductBinary(tt.args.strs); got != tt.want {
				t.Errorf("maxProductBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

type generator05 struct {
	src *rand.Rand
}

func (g *generator05) nextInt(max int) int {
	return g.src.Intn(max)
}

func (g *generator05) nextString(maxLen int) string {
	var sb strings.Builder
	sb.Grow(maxLen)

	strLen := g.nextInt(maxLen)
	for i := 0; i < strLen; i++ {
		sb.WriteByte('a' + byte(g.nextInt(26)))
	}

	return sb.String()
}

func testSet05(seed, num int) []string {
	gen := generator05{
		src: rand.New(rand.NewSource(int64(seed))),
	}

	strs := make([]string, num)
	for i := range strs {
		strs[i] = gen.nextString(num)
	}

	return strs
}

func BenchmarkMaxProduct(b *testing.B) {
	tests := []struct {
		name string
		f    func(strs []string) int
	}{
		{name: "BruteForce", f: maxProductBruteForce},
		{name: "HashTable", f: maxProductHashTable},
		{name: "Binary", f: maxProductBinary},
	}
	for k := 10; k <= 10000; k *= 10 {
		strs := testSet05(300, k)
		b.ResetTimer()
		for _, tt := range tests {
			b.Run(fmt.Sprintf("%s-%d", tt.name, k), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tt.f(strs)
				}
			})
		}
	}
}
