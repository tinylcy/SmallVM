package rtarea

type Frame struct {
	thread         *Thread
	localVariables *LocalVariables
	operandStack   *OperandStack
	code           []byte
	next           *Frame
}

func NewFrame(thread *Thread, maxLocals uint, maxStack uint, code []byte) *Frame {
	frame := &Frame{}
	frame.thread = thread
	frame.localVariables = NewLocalVariables(maxLocals)
	frame.operandStack = NewOperandStack(maxStack)
	frame.code = code
	frame.next = nil
	return frame
}

func (self *Frame) LocalVariables() *LocalVariables {
	return self.localVariables
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) CurrentThread() *Thread {
	return self.thread
}

func (self *Frame) Code() []byte {
	return self.code
}
