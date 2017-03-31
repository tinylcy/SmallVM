package constant

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type BIPUSH struct {
	value int8
}

func (self *BIPUSH) FetchOperands(reader *common.ByteCodeReader) {
	self.value = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtarea.Frame) {
	int32Val := int32(self.value)
	frame.OperandStack().PushInt(int32Val)
}
