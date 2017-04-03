package load

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type ALOAD struct {
	index uint
}

type ALOAD_0 struct {
}

type ALOAD_1 struct {
}

type ALOAD_2 struct {
}

type ALOAD_3 struct {
}

func (self *ALOAD) FetchOperands(reader *common.ByteCodeReader) {
	self.index = uint(reader.ReadUInt8())
}

func (self *ALOAD) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetReference(self.index)
	frame.OperandStack().PushReference(value)
}

func (self *ALOAD) SetIndex(index uint) {
	self.index = index
}
func (self *ALOAD_0) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ALOAD_0) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetReference(0)
	frame.OperandStack().PushReference(value)
}

func (self *ALOAD_1) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ALOAD_1) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetReference(1)
	frame.OperandStack().PushReference(value)
}

func (self *ALOAD_2) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ALOAD_2) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetReference(2)
	frame.OperandStack().PushReference(value)
}

func (self *ALOAD_3) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ALOAD_3) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetReference(3)
	frame.OperandStack().PushReference(value)
}
