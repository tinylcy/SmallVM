package math

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Add two integers
type IADD struct {
}

// Add two long integers
type LADD struct {
}

// Add two floats
type FADD struct {
}

// Add two doubles
type DADD struct {
}

func (self *IADD) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *IADD) Execute(frame rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopInt()
	value2 := stack.PopInt()
	stack.PushInt(value1 + value2)
}

func (self *LADD) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LADD) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopLong()
	value2 := stack.PopLong()
	stack.PushLong(value1 + value2)
}

func (self *FADD) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FADD) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopFloat()
	value2 := stack.PopFloat()
	stack.PushFloat(value1 + value2)
}

func (self *DADD) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DADD) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopDouble()
	value2 := stack.PopDouble()
	stack.PushDouble(value1 + value2)
}
