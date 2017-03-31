package store

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type ISTORE struct {
	index uint
}

type ISTORE_0 struct {
}

type ISTORE_1 struct {
}

type ISTORE_2 struct {
}

type ISTORE_3 struct {
}

func (self *ISTORE) FetchOperands(reader *common.ByteCodeReader) {
	self.index = uint(reader.ReadUInt8())
}

func (self *ISTORE) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopInt()
	frame.LocalVariables().SetInt(self.index, value)
}

func (self *ISTORE_0) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ISTORE_0) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopInt()
	frame.LocalVariables().SetInt(0, value)
}

func (self *ISTORE_1) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ISTORE_1) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopInt()
	frame.LocalVariables().SetInt(1, value)
}

func (self *ISTORE_2) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ISTORE_2) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopInt()
	frame.LocalVariables().SetInt(2, value)
}

func (self *ISTORE_3) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ISTORE_3) Execute(frame *rtarea.Frame) {
	value := frame.OperandStack().PopInt()
	frame.LocalVariables().SetInt(3, value)
}
