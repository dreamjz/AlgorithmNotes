package ofgraph

import (
	"testing"
)

func TestIsBipartiteBFS(t *testing.T) {
	type args struct {
		graph [][]int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "#01", args: args{graph: [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}}, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBipartiteBFS(tt.args.graph); got != tt.want {
				t.Errorf("isBipartite() = %v, want %v", got, tt.want)
			}
		})
	}
}
