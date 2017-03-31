package constant

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type ICONST_M1 struct {
}

type ICONST_0 struct {
}

type ICONST_1 struct {
}

type ICONST_2 struct {
}

type ICONST_3 struct {
}

type ICONST_4 struct {
}

type ICONST_5 struct {
}

func (self *ICONST_M1) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ICONST_M1) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PushInt(int32(-1))
}

func (self *ICONST_0) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ICONST_0) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PushInt(int32(0))
}

func (self *ICONST_1) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ICONST_1) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PushInt(int32(1))
}

func (self *ICONST_2) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ICONST_2) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PushInt(int32(2))
}

func (self *ICONST_3) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ICONST_3) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PushInt(int32(3))
}

func (self *ICONST_4) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ICONST_4) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PushInt(int32(4))
}

func (self *ICONST_5) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ICONST_5) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PushInt(int32(5))
}
