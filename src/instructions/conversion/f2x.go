package conversion

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Convert float to integer
type F2I struct {
}

// Convert float to long integer
type F2L struct {
}

// Convert float to double
type F2D struct {
}

func (self *F2I) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *F2I) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopFloat()
	stack.PushInt(int32(value))
}

func (self *F2L) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *F2L) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopFloat()
	stack.PushLong(int64(value))
}

func (self *F2D) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *F2D) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopFloat()
	stack.PushDouble(float64(value))
}
