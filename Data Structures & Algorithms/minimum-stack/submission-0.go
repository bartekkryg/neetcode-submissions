type MinStack struct {
	values []int
	min []int
}

func Constructor() MinStack {
	return MinStack{
		values: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.values = append(this.values, val)
	minVal := val
	if len(this.min) > 0 {
		if this.min[len(this.min)-1] < val {
			minVal = this.min[len(this.min)-1]
		}
	}
	this.min = append(this.min, minVal)
}

func (this *MinStack) Pop() {
	n := len(this.values)
	if n > 0 {
		this.values = this.values[:n-1]
	}
	k := len(this.min)
	if k > 0 {
		this.min = this.min[:k-1]
	}
}

func (this *MinStack) Top() int {
	var r int
	n := len(this.values)
	if n > 0 {
		r = this.values[n-1]
	}
	return r
}

func (this *MinStack) GetMin() int {
	if len(this.min) > 0 {
		return this.min[len(this.min)-1]
	}
	return 0 
}
