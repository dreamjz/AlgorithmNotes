package ordastring

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "empty string", args: args{s: ""}, want: 0},
		{name: "#01", args: args{s: "abcabcbb"}, want: 3},
		{name: "#02", args: args{s: "bbbbb"}, want: 1},
		{name: "#03", args: args{s: "pwwkew"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstringWithCountDup(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "empty string", args: args{s: ""}, want: 0},
		{name: "#01", args: args{s: "abcabcbb"}, want: 3},
		{name: "#02", args: args{s: "bbbbb"}, want: 1},
		{name: "#03", args: args{s: "pwwkew"}, want: 3},
		{name: "#04", args: args{s: "nfpdmpi"}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringWithCountDup(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringWithCountDup() = %v, want %v", got, tt.want)
			}
		})
	}
}
