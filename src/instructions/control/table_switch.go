package control

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Jump according to a table
type TABLESWITCH struct {
	basePC        int // thread's pc value when TABLESWITCH starts to execute
	defaultOffset int32
	low           int32
	high          int32
	table         []int32
}

func (self *TABLESWITCH) FetchOperands(reader *common.ByteCodeReader) {
	self.basePC = reader.CurrentPC() - 1
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	tableLength := self.high - self.low + 1 // There must be <high>-<low>+1 labels given in the table.
	self.table = make([]int32, tableLength)
	var i int32
	for i = 0; i < tableLength; i++ {
		self.table[i] = reader.ReadInt32()
	}
}

// If value is less than <low>, or if it is greater than <high>, execution branches to the address (pc + default_offset).
// If value is in the range <low> to <high>, execution branches to the i'th entry in the table,
// where i is (val - <low>) and the first entry in the table is index 0.
func (self *TABLESWITCH) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopInt()
	pc := self.basePC
	var nextPC int
	if value < self.low || value > self.high {
		nextPC = pc + int(self.defaultOffset)
	} else {
		nextPC = pc + int(self.table[value-self.low])
	}
	frame.CurrentThread().SetPC(nextPC)
}
