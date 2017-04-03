package rtarea

const MAX_FRAME_DEPTH = 1024

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	thread := &Thread{}
	thread.pc = 0
	thread.stack = NewStack(MAX_FRAME_DEPTH)
	return thread
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(value int) {
	self.pc = value
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.Push(frame)
}

func (self *Thread) PopFrame() *Frame {
	frame := self.stack.Pop()
	return frame
}

func (self *Thread) CurrentFrame() *Frame {
	frame := self.stack.Peek()
	return frame
}
