package constant

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type ACONST_NULL struct {
}

func (self *ACONST_NULL) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *ACONST_NULL) Execute(frame rtarea.Frame) {
	frame.OperandStack().PushReference(nil)
}
