package load

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type DLOAD struct {
	index uint
}

type DLOAD_0 struct {
}

type DLOAD_1 struct {
}

type DLOAD_2 struct {
}

type DLOAD_3 struct {
}

func (self *DLOAD) FetchOperands(reader *common.ByteCodeReader) {
	self.index = uint(reader.ReadUInt8())
}

func (self *DLOAD) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetDouble(self.index)
	frame.OperandStack().PushDouble(value)
}

func (self *DLOAD) SetIndex(index uint) {
	self.index = index
}

func (self *DLOAD_0) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DLOAD_0) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetDouble(0)
	frame.OperandStack().PushDouble(value)
}

func (self *DLOAD_1) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DLOAD_1) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetDouble(1)
	frame.OperandStack().PushDouble(value)
}

func (self *DLOAD_2) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DLOAD_2) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetDouble(2)
	frame.OperandStack().PushDouble(value)
}

func (self *DLOAD_3) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DLOAD_3) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetDouble(3)
	frame.OperandStack().PushDouble(value)
}
