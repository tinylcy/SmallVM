package math

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Subtract two integers
type ISUB struct {
}

// Subtract two longs
type LSUB struct {
}

// Subtract two floats
type FSUB struct {
}

// Subtract two doubles
type DSUB struct {
}

func (self *ISUB) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ISUB) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopInt()
	value2 := stack.PopInt()
	stack.PushInt(value2 - value1)
}

func (self *LSUB) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LSUB) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopLong()
	value2 := stack.PopLong()
	stack.PushLong(value2 - value1)
}

func (self *FSUB) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FSUB) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopFloat()
	value2 := stack.PopFloat()
	stack.PushFloat(value2 - value1)
}

func (self *DSUB) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DSUB) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopDouble()
	value2 := stack.PopDouble()
	stack.PushDouble(value2 - value1)
}
