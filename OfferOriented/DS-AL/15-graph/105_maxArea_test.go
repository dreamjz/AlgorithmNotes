package ofgraph

import "testing"

func TestAreaOfIslandBFS(t *testing.T) {
	type args struct {
		grid [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "#01", args: args{grid: [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}}, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaOfIslandBFS(tt.args.grid); got != tt.want {
				t.Errorf("maxAreaOfIslandBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
