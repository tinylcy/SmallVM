package control

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type LOOKUPSWITCH struct {
	basePC        int // thread's pc value when LOOKUPSWITCH starts to execute
	defaultOffset int32
	labelsCount   int32
	table         []int32
}

func (self *LOOKUPSWITCH) FetchOperands(reader *common.ByteCodeReader) {
	self.basePC = reader.CurrentPC() - 1
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.labelsCount = reader.ReadInt32()
	self.table = make([]int32, self.labelsCount*2)
	var i int32
	for i = 0; i < self.labelsCount*2; i++ {
		self.table[i] = reader.ReadInt32()
	}
}

// First, an int, item, is taken from the top of the stack.
// Then, lookupswitch searches the table looking for an entry whose <key> field matches item.
// If a match is found, execution branches to the address of the corresponding <label>.
// If no match is found, execution branches to <labelDefault>.
func (self *LOOKUPSWITCH) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopInt()
	var i int32
	pc := self.basePC
	var nextPC int
	for i = 0; i < self.labelsCount*2; i += 2 {
		key := self.table[i]
		if key == value {
			nextPC = pc + int(self.table[i+1])
			frame.CurrentThread().SetPC(nextPC)
			return
		}
	}
	nextPC = pc + int(self.defaultOffset)
	frame.CurrentThread().SetPC(nextPC)
}
