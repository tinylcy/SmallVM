package store

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type ASTORE struct {
	index uint
}

type ASTORE_0 struct {
}

type ASTORE_1 struct {
}

type ASTORE_2 struct {
}

type ASTORE_3 struct {
}

func (self *ASTORE) FetchOperands(reader *common.ByteCodeReader) {
	self.index = uint(reader.ReadUInt8())
}

func (self *ASTORE) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopReference()
	frame.LocalVariables().SetReference(self.index, value)
}

func (self *ASTORE) SetIndex(index uint) {
	self.index = index
}

func (self *ASTORE_0) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ASTORE_0) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopReference()
	frame.LocalVariables().SetReference(0, value)
}

func (self *ASTORE_1) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ASTORE_1) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopReference()
	frame.LocalVariables().SetReference(1, value)
}

func (self *ASTORE_2) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ASTORE_2) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopReference()
	frame.LocalVariables().SetReference(2, value)
}

func (self *ASTORE_3) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ASTORE_3) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopReference()
	frame.LocalVariables().SetReference(3, value)
}
