package ofqueue

type MovingAverage struct {
	queue []int
	size  int
	sum   int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		queue: make([]int, 0, size),
		size:  size,
	}
}

func (ma *MovingAverage) Next(val int) float64 {
	var average float64
	// 滑动窗口未满
	if len(ma.queue) < ma.size {
		ma.queue = append(ma.queue, val)
		ma.sum += val
		average = float64(ma.sum) / float64(len(ma.queue))
		return average
	}
	// 滑动窗口已满
	// 队尾出队, 新元素入队
	tail := ma.queue[0]
	ma.queue = ma.queue[1:]
	ma.queue = append(ma.queue, val)
	// 平均值
	ma.sum = ma.sum + val - tail
	average = float64(ma.sum) / float64(len(ma.queue))

	return average
}
