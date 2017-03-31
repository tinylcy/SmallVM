package constant

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type FCONST_0 struct {
}

type FCONST_1 struct {
}

type FCONST_2 struct {
}

func (self *FCONST_0) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FCONST_0) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PushFloat(float32(0))
}

func (self *FCONST_1) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FCONST_1) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PushFloat(float32(1))
}

func (self *FCONST_2) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *FCONST_2) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PushFloat(float32(2))
}
