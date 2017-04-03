package store

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type LSTORE struct {
	index uint
}

type LSTORE_0 struct {
}

type LSTORE_1 struct {
}

type LSTORE_2 struct {
}

type LSTORE_3 struct {
}

func (self *LSTORE) FetchOperands(reader *common.ByteCodeReader) {
	self.index = uint(reader.ReadUInt8())
}

func (self *LSTORE) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopLong()
	frame.LocalVariables().SetLong(self.index, value)
}

func (self *LSTORE) SetIndex(index uint) {
	self.index = index
}

func (self *LSTORE_0) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LSTORE_0) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopLong()
	frame.LocalVariables().SetLong(0, value)
}

func (self *LSTORE_1) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LSTORE_1) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopLong()
	frame.LocalVariables().SetLong(1, value)
}

func (self *LSTORE_2) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LSTORE_2) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopLong()
	frame.LocalVariables().SetLong(2, value)
}

func (self *LSTORE_3) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LSTORE_3) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopLong()
	frame.LocalVariables().SetLong(3, value)
}
