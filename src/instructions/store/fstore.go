package store

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type FSTORE struct {
	index uint
}

type FSTORE_0 struct {
}

type FSTORE_1 struct {
}

type FSTORE_2 struct {
}

type FSTORE_3 struct {
}

func (self *FSTORE) FetchOperands(reader *common.ByteCodeReader) {
	self.index = uint(reader.ReadUInt8())
}

func (self *FSTORE) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopFloat()
	frame.LocalVariables().SetFloat(self.index, value)
}

func (self *FSTORE) SetIndex(index uint) {
	self.index = index
}

func (self *FSTORE_0) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FSTORE_0) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopFloat()
	frame.LocalVariables().SetFloat(0, value)
}

func (self *FSTORE_1) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FSTORE_1) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopFloat()
	frame.LocalVariables().SetFloat(1, value)
}

func (self *FSTORE_2) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FSTORE_2) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopFloat()
	frame.LocalVariables().SetFloat(2, value)
}

func (self *FSTORE_3) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FSTORE_3) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopFloat()
	frame.LocalVariables().SetFloat(3, value)
}
