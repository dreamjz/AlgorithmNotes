package ofqueue

type RecentCounter struct {
	queue []int
	size  int
}

func Constructor42() RecentCounter {
	return RecentCounter{
		queue: make([]int, 0),
		size:  3000,
	}
}

func (rc *RecentCounter) Ping(t int) int {
	// 空窗口
	if len(rc.queue) == 0 {
		rc.queue = append(rc.queue, t)
		return len(rc.queue)
	}

	// 将超过3000毫秒的请求出队
	for len(rc.queue) > 0 && rc.queue[0]+rc.size < t {
		rc.queue = rc.queue[1:]
	}
	rc.queue = append(rc.queue, t)
	// 返回个数
	return len(rc.queue)
}
