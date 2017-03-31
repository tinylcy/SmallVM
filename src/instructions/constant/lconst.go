package constant

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type LCONST_0 struct {
}

type LCONST_1 struct {
}

func (self *LCONST_0) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LCONST_0) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PushLong(int64(0))
}

func (self *LCONST_1) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *LCONST_1) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PushLong(int64(1))
}
