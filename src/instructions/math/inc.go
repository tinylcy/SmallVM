package math

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// For local variable numbers <index> in the range 0-255, and values of <increment> in the range -128 to 127
type IINC struct {
	index     uint
	increment int32
}

func (self *IINC) FetchOperands(reader *common.ByteCodeReader) {
	self.index = uint(reader.ReadUInt8())     // 0-255
	self.increment = int32(reader.ReadInt8()) // -128-127
}

// Increments the int held in the local variable <index> by <increment>.
func (self *IINC) Execute(frame *rtarea.Frame) {
	localVariabls := frame.LocalVariables()
	value := localVariabls.GetInt(self.index)
	localVariabls.SetInt(self.index, value+self.increment)
}

func (self *IINC) SetIndex(index uint) {
	self.index = index
}

func (self *IINC) SetIncrement(increment int32) {
	self.increment = increment
}
