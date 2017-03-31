package math

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Divide an integer by another integer
type IDIV struct {
}

// Divide a long integer
type LDIV struct {
}

// Divide two floats
type FDIV struct {
}

// Divide two doubles
type DDIV struct {
}

func (self *IDIV) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *IDIV) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopInt()
	value2 := stack.PopInt()
	if value1 == 0 {
		panic("java.lang.AithemicException: / by zero.")
	}
	stack.PushInt(value2 / value1)
}

func (self *LDIV) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LDIV) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopLong()
	value2 := stack.PopLong()
	if value1 == 0 {
		panic("java.lang.AithemicException: / by zero.")
	}
	stack.PushLong(value2 / value1)
}

func (self *FDIV) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FDIV) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopFloat()
	value2 := stack.PopFloat()
	if value1 == float32(0) {
		panic("java.lang.AithemicException: / by zero.")
	}
	stack.PushFloat(value2 / value1)
}

func (self *DDIV) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DDIV) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopDouble()
	value2 := stack.PopDouble()
	if value1 == float64(0) {
		panic("java.lang.AithemicException: / by zero.")
	}
	stack.PushDouble(value2 / value1)
}
