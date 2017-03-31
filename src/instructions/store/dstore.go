package store

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type DSTORE struct {
	index uint
}

type DSTORE_0 struct {
}

type DSTORE_1 struct {
}

type DSTORE_2 struct {
}

type DSTORE_3 struct {
}

func (self *DSTORE) FetchOperands(reader *common.ByteCodeReader) {
	self.index = uint(reader.ReadUInt8())
}

func (self *DSTORE) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopDouble()
	frame.LocalVariables().SetDouble(self.index, value)
}

func (self *DSTORE_0) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DSTORE_0) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopDouble()
	frame.LocalVariables().SetDouble(0, value)
}

func (self *DSTORE_1) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DSTORE_1) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopDouble()
	frame.LocalVariables().SetDouble(1, value)
}

func (self *DSTORE_2) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DSTORE_2) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopDouble()
	frame.LocalVariables().SetDouble(2, value)
}

func (self *DSTORE_3) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DSTORE_3) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopDouble()
	frame.LocalVariables().SetDouble(3, value)
}
