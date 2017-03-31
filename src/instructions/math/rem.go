package math

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
	"math"
)

// Remainder of two integers
type IREM struct {
}

// Remainder of two longs
type LREM struct {
}

// Remainder of two floats
type FREM struct {
}

// Remainder of two doubles
type DREM struct {
}

func (self *IREM) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *IREM) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopInt()
	value2 := stack.PopInt()
	if value1 == 0 {
		panic("java.lang.AithemicException: % by zero.")
	}
	stack.PushInt(value2 % value1)
}

func (self *LREM) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LREM) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopLong()
	value2 := stack.PopLong()
	if value1 == 0 {
		panic("java.lang.AithemicException: % by zero.")
	}
	stack.PushLong(value2 % value1)
}

func (self *FREM) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FREM) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := float64(stack.PopFloat())
	value2 := float64(stack.PopFloat())
	if value1 == float64(0) {
		panic("java.lang.AithemicException: % by zero.")
	}
	result := math.Mod(value2, value1)
	stack.PushFloat(float32(result))
}

func (self *DREM) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DREM) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopDouble()
	value2 := stack.PopDouble()
	if value1 == float64(0) {
		panic("java.lang.AithemicException: % by zero.")
	}
	stack.PushDouble(math.Mod(value2, value1))
}
