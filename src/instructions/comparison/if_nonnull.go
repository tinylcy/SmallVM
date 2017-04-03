package comparison

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Jump if not null
type IFNONNULL struct {
	offset int16
}

func (self *IFNONNULL) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

func (self *IFNONNULL) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopReference()
	if value != nil {
		pc := frame.CurrentThread().PC() - 3
		nextPC := pc + int(self.offset)
		frame.CurrentThread().SetPC(nextPC)
	}
}
