package load

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type LLOAD struct {
	index uint
}

type LLOAD_0 struct {
}

type LLOAD_1 struct {
}

type LLOAD_2 struct {
}

type LLOAD_3 struct {
}

func (self *LLOAD) FetchOperands(reader *common.ByteCodeReader) {
	self.index = uint(reader.ReadUInt8())
}

func (self *LLOAD) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetLong(self.index)
	frame.OperandStack().PushLong(value)
}

func (self *LLOAD_0) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LLOAD_0) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetLong(0)
	frame.OperandStack().PushLong(value)
}

func (self *LLOAD_1) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LLOAD_1) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetLong(1)
	frame.OperandStack().PushLong(value)
}

func (self *LLOAD_2) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LLOAD_2) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetLong(2)
	frame.OperandStack().PushLong(value)
}

func (self *LLOAD_3) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LLOAD_3) Execute(frame *rtarea.Frame) {
	value := frame.LocalVariables().GetLong(3)
	frame.OperandStack().PushLong(value)
}
