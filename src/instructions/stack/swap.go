package stack

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type SWAP struct {
}

func (self *SWAP) FetchOperands(reader *common.ByteCodeReader) {
}

// Swap top two stack slots
func (self *SWAP) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
