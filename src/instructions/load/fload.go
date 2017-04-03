package load

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type FLOAD struct {
	index uint
}

type FLOAD_0 struct {
}

type FLOAD_1 struct {
}

type FLOAD_2 struct {
}

type FLOAD_3 struct {
}

func (self *FLOAD) FetchOperands(reader *common.ByteCodeReader) {
	self.index = uint(reader.ReadUInt8())
}

func (self *FLOAD) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetFloat(self.index)
	frame.OperandStack().PushFloat(value)
}

func (self *FLOAD) SetIndex(index uint) {
	self.index = index
}

func (self *FLOAD_0) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FLOAD_0) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetFloat(0)
	frame.OperandStack().PushFloat(value)
}

func (self *FLOAD_1) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FLOAD_1) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetFloat(1)
	frame.OperandStack().PushFloat(value)
}

func (self *FLOAD_2) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FLOAD_2) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetFloat(2)
	frame.OperandStack().PushFloat(value)
}

func (self *FLOAD_3) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FLOAD_3) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetFloat(3)
	frame.OperandStack().PushFloat(value)
}
