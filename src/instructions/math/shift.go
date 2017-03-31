package math

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Integer shift left
type ISHL struct {
}

// Long integer shift left
type LSHL struct {
}

// Integer arithmetic shift right
type ISHR struct {
}

// Long integer arithmetic shift right
type LSHR struct {
}

// Integer logical shift right
type IUSHR struct {
}

// Long integer logical shift right
type LUSHR struct {
}

func (self *ISHL) FetchOperands(reader *common.ByteCodeReader) {
}

//  Shifts value2 left by the amount indicated in the five low bits of value1.
func (self *ISHL) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := uint32(stack.PopInt()) & 0x1F
	value2 := stack.PopInt()
	result := value2 << value1
	stack.PushInt(result)
}

func (self *LSHL) FetchOperands(reader *common.ByteCodeReader) {
}

// Shifts value2 (the long integer) left by the amount indicated in the low six bits of value1 (an int).
func (self *LSHL) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := uint32(stack.PopInt()) & 0x1F
	value2 := stack.PopLong()
	result := value2 << value1
	stack.PushLong(result)
}

func (self *ISHR) FetchOperands(reader *common.ByteCodeReader) {
}

// Shifts value1 right by the amount indicated in the five low bits of value2.
func (self *ISHR) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopInt()
	value2 := uint32(stack.PopInt()) & 0x1F
	result := value1 >> value2
	stack.PushInt(result)
}

func (self *LSHR) FetchOperands(reader *common.ByteCodeReader) {
}

// Shifts value2 (the long integer) right by the amount indicated in the low six bits of value1 (an int)
func (self *LSHR) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := uint32(stack.PopInt()) & 0x1F
	value2 := stack.PopLong()
	result := value2 >> value1
	stack.PushLong(result)
}

func (self *IUSHR) FetchOperands(reader *common.ByteCodeReader) {
}

// Shifts value1 right by the amount indicated in the five low bits of value2.
func (self *IUSHR) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := uint32(stack.PopInt())
	value2 := uint32(stack.PopInt()) & 0x1F
	result := int32(value1 >> value2)
	stack.PushInt(result)
}

func (self *LUSHR) FetchOperands(reader *common.ByteCodeReader) {
}

// Shifts value2 (the long integer) right by the amount indicated in the low six bits of value1 (an int).
func (self *LUSHR) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := uint32(stack.PopInt()) & 0x3F
	value2 := uint64(stack.PopLong())
	result := int64(value2 >> value1)
	stack.PushLong(result)
}
