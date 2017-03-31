package math

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Negate an integer
type INEG struct {
}

// Negate a long
type LNEG struct {
}

// Negate a float
type FNEG struct {
}

// Negate a double
type DNEG struct {
}

func (self *INEG) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *INEG) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopInt()
	stack.PushInt(-1 * value)
}

func (self *LNEG) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LNEG) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopLong()
	stack.PushLong(-1 * value)
}

func (self *FNEG) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FNEG) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopFloat()
	stack.PushFloat(-1 * value)
}

func (self *DNEG) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DNEG) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopDouble()
	stack.PushDouble(-1 * value)
}
