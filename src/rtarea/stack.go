package rtarea

type Stack struct {
	top     *Frame
	size    uint
	maxSize uint
}

func NewStack(maxSize uint) *Stack {
	if maxSize > 0 {
		stack := &Stack{}
		stack.top = nil
		stack.size = 0
		stack.maxSize = maxSize
		return stack
	}
	return nil
}

func (self *Stack) Push(frame *Frame) {
	if self.size == self.maxSize {
		panic("java.lang.StackOverflowError!")
	}
	if self.top == nil {
		self.top = frame
	} else {
		frame.next = self.top
		self.top = frame
	}
	self.size++
}

func (self *Stack) Pop() *Frame {
	if self.size == 0 {
		panic("JVM stack is empty!")
	}
	topElem := self.top
	self.top = topElem.next
	topElem.next = nil
	self.size--

	return topElem
}

func (self *Stack) Peek() *Frame {
	if self.size == 0 {
		panic("JVM stack is empty!")
	}
	return self.top
}
