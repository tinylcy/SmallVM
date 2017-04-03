package comparison

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Jump if two object references are equal
type IF_ACMPEQ struct {
	offset int16
}

// Jump if two object references are not equal
type IF_ACMPNE struct {
	offset int16
}

func (self *IF_ACMPEQ) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

func (self *IF_ACMPEQ) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopReference()
	value2 := stack.PopReference()
	if value1 == value2 {
		pc := frame.CurrentThread().PC() - 3
		nextPC := pc + int(self.offset)
		frame.CurrentThread().SetPC(nextPC)
	}
}

func (self *IF_ACMPNE) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

func (self *IF_ACMPNE) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopReference()
	value2 := stack.PopReference()
	if value1 != value2 {
		pc := frame.CurrentThread().PC() - 3
		nextPC := pc + int(self.offset)
		frame.CurrentThread().SetPC(nextPC)
	}
}
