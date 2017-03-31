package stack

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type POP struct {
}

type POP2 struct {
}

func (self *POP) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *POP) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PopSlot()
}

func (self *POP2) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *POP2) Execute(frame *rtarea.Frame) {
	frame.OperandStack().PopSlot()
	frame.OperandStack().PopSlot()
}
