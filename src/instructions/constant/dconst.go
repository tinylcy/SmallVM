package constant

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type DCONST_0 struct {
}

type DCONST_1 struct {
}

func (self *DCONST_0) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DCONST_0) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PushDouble(float64(0))
}

func (self *DCONST_1) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DCONST_1) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PushDouble(float64(1))
}
