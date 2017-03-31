package conversion

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Convert integer to long integer
type I2L struct {
}

// Convert integer to float
type I2F struct {
}

// Convert integer to double
type I2D struct {
}

// Convert integer to byte
type I2B struct {
}

// Convert integer to character
type I2C struct {
}

// Convert integer to short integer
type I2S struct {
}

func (self *I2L) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *I2L) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := int64(stack.PopInt())
	stack.PushLong(value)
}

func (self *I2F) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *I2F) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := float32(stack.PopInt())
	stack.PushFloat(value)
}

func (self *I2D) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *I2D) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := float64(stack.PopInt())
	stack.PushDouble(value)
}

func (self *I2B) FetchOperands(reader *common.ByteCodeReader) {
}

// Converts an integer to a signed byte.
func (self *I2B) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := int8(stack.PopInt())
	stack.PushInt(int32(value))
}

func (self *I2C) FetchOperands(reader *common.ByteCodeReader) {
}

// Converts an integer to a 16-bit unsigned char.
func (self *I2C) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := uint16(stack.PopInt())
	stack.PushInt(int32(value))
}

func (self *I2S) FetchOperands(reader *common.ByteCodeReader) {
}

// Converts an integer to a signed short.
func (self *I2S) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := int16(stack.PopInt())
	stack.PushInt(int32(value))
}
