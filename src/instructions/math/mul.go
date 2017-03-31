package math

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Multiply two integers
type IMUL struct {
}

// Multiply two longs
type LMUL struct {
}

// Multiply two floats
type FMUL struct {
}

// Multiply two longs
type DMUL struct {
}

func (self *IMUL) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *IMUL) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopInt()
	value2 := stack.PopInt()
	stack.PushInt(value1 * value2)
}

func (self *LMUL) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LMUL) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopLong()
	value2 := stack.PopLong()
	stack.PushLong(value1 * value2)
}

func (self *FMUL) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FMUL) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopFloat()
	value2 := stack.PopFloat()
	stack.PushFloat(value1 * value2)
}

func (self *DMUL) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DMUL) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopDouble()
	value2 := stack.PopDouble()
	stack.PushDouble(value1 * value2)
}
