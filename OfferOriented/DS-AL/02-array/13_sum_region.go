package dsalarr

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m == 0 {
		return NumMatrix{}
	}
	n := len(matrix[0])
	sums := make([][]int, m+1)
	sums[0] = make([]int, n+1)
	for r := range matrix {
		sums[r+1] = make([]int, n+1)
		for c := range matrix[r] {
			sums[r+1][c+1] = sums[r+1][c] + sums[r][c+1] - sums[r][c] + matrix[r][c]
		}
	}

	return NumMatrix{sums: sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sums := this.sums
	return sums[row2+1][col2+1] - sums[row2+1][col1] - sums[row1][col2+1] + sums[row1][col1]
}
