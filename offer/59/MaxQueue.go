package MaxQueue

type MaxQueue struct {
	nums  []int
	queue []int
}

func Constructor() MaxQueue {
	nums := make([]int, 0, 100001)
	queue := make([]int, 0, 100001)
	return MaxQueue{nums, queue}
}

func (this *MaxQueue) Max_value() int {
	if len(this.queue) > 0 {
		return this.queue[0]
	}
	return -1
}

func (this *MaxQueue) Push_back(value int) {
	this.nums = append(this.nums, value)
	for len(this.queue) > 0 && value > this.queue[len(this.queue)-1] {
		this.queue = this.queue[:len(this.queue)-1]
	}
	this.queue = append(this.queue, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.nums) == 0 {
		return -1
	}
	var num int
	num, this.nums = this.nums[0], this.nums[1:]
	if num == this.queue[0] {
		this.queue = this.queue[1:]
	}
	return num
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
