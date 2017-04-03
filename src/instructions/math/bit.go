package math

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Integer bitwise and
type IAND struct {
}

// Long integer bitwise and
type LAND struct {
}

// Integer bitwise or
type IOR struct {
}

// Long integer bitwise or
type LOR struct {
}

// Integer bitwise exclusive or
type IXOR struct {
}

// Long integer bitwise exclusive or
type LXOR struct {
}

func (self *IAND) FetchOperands(reader *common.ByteCodeReader) {
}

// Computes the bitwise and of value1 and value2 (which must be ints).
func (self *IAND) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopInt()
	value2 := stack.PopInt()
	result := value1 & value2
	stack.PushInt(result)
}

func (self *LAND) FetchOperands(reader *common.ByteCodeReader) {
}

// Pops two long integers off the stack. Computes the bitwise and of value1 and value2.
func (self *LAND) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopLong()
	value2 := stack.PopLong()
	result := value1 & value2
	stack.PushLong(result)
}

func (self *IOR) FetchOperands(reader *common.ByteCodeReader) {
}

// Computes the bitwise or of value1 and value2 (which must be ints).
func (self *IOR) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopInt()
	value2 := stack.PopInt()
	result := value1 | value2
	stack.PushInt(result)
}

func (self *LOR) FetchOperands(reader *common.ByteCodeReader) {
}

// Computes the bitwise or of value1 and value2.
func (self *LOR) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopLong()
	value2 := stack.PopLong()
	result := value1 | value2
	stack.PushLong(result)
}

func (self *IXOR) FetchOperands(reader *common.ByteCodeReader) {
}

// Computes the bitwise exclusive or of value1 and value2.
func (self *IXOR) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopInt()
	value2 := stack.PopInt()
	result := value1 ^ value2
	stack.PushInt(result)
}

func (self *LXOR) FetchOperands(reader *common.ByteCodeReader) {
}

// Computes the bitwise exclusive or of value1 and value2.
func (self *LXOR) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopLong()
	value2 := stack.PopLong()
	result := value1 ^ value2
	stack.PushLong(result)
}
