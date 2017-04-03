package load

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type ILOAD struct {
	index uint
}

type ILOAD_0 struct {
}

type ILOAD_1 struct {
}

type ILOAD_2 struct {
}

type ILOAD_3 struct {
}

func (self *ILOAD) FetchOperands(reader *common.ByteCodeReader) {
	value := uint(reader.ReadUInt8())
	self.index = value
}

func (self *ILOAD) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetInt(self.index)
	frame.OperandStack().PushInt(value)
}

func (self *ILOAD) SetIndex(index uint) {
	self.index = index
}

func (self *ILOAD_0) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ILOAD_0) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetInt(0)
	frame.OperandStack().PushInt(value)
}

func (self *ILOAD_1) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ILOAD_1) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetInt(1)
	frame.OperandStack().PushInt(value)
}

func (self *ILOAD_2) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ILOAD_2) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetInt(2)
	frame.OperandStack().PushInt(value)
}

func (self *ILOAD_3) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ILOAD_3) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetInt(3)
	frame.OperandStack().PushInt(value)
}
