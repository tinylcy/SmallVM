package constant

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type SIPUSH struct {
	value int16
}

func (self *SIPUSH) FetchOperands(reader *common.ByteCodeReader) {
	self.value = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtarea.Frame) {
	int32Val := int32(self.value)
	frame.OperandStack().PushInt(int32Val)
}
