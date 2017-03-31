package conversion

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Convert double to integer
type D2I struct {
}

// Convert double to long integer
type D2L struct {
}

// Convert double to float
type D2F struct {
}

func (self *D2I) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *D2I) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopDouble()
	stack.PushInt(int32(value))
}

func (self *D2L) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *D2L) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopDouble()
	stack.PushLong(int64(value))
}

func (self *D2F) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *D2F) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopDouble()
	stack.PushFloat(float32(value))
}
