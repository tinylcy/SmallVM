package conversion

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Long to integer conversion
type L2I struct {
}

// Convert long to float
type L2F struct {
}

// Convert long to double
type L2D struct {
}

func (self *L2I) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *L2I) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := int32(stack.PopLong())
	stack.PushInt(value)
}

func (self *L2F) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *L2F) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := float32(stack.PopLong())
	stack.PushFloat(value)
}

func (self *L2D) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *L2D) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := float64(stack.PopLong())
	stack.PushDouble(value)
}
