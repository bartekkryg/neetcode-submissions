func calPoints(operations []string) int {
	opStack := New(make([]string, 0))
	for _, op := range operations {
		if isNum(op) {
			opStack.Push(op)
		} else if op == "C" {
			opStack.Pop()
		} else if op == "D" {
			x := opStack.Pop()
			xInt, _ := strconv.Atoi(x)
			opStack.Push(x)
			opStack.Push(strconv.Itoa(2*xInt))
		} else if op == "+" {
			x := opStack.Pop()
			y := opStack.Pop()
			xInt, _ := strconv.Atoi(x)
			yInt, _ := strconv.Atoi(y)
			opStack.Push(y)
			opStack.Push(x)
			opStack.Push(strconv.Itoa(xInt+yInt))
		}
	}
	return opStack.Sum()
}

type stack struct {
	values []string
}

func New(arr []string) *stack {
	return &stack{
		values: arr,
	}

}


func (s *stack) Pop() string {
	if len(s.values) > 0 {
		el := s.values[len(s.values)-1]
		s.values = s.values[:len(s.values)-1]
		return el
	}
	return ""
}

func (s *stack) Len() int {
	return len(s.values)
}

func (s *stack) Push(el string) {
	s.values = append(s.values, el)
}

func (s *stack) Sum() int {
	total := 0
	for _, x := range s.values {
		xInt, _ := strconv.Atoi(x)
		total += xInt
	}
	return total
}

func isNum(el string) bool {
	return el != "C" && el != "D" && el != "+"
}