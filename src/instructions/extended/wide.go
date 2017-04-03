package extended

import (
	"SmallVM/instructions/common"
	"SmallVM/instructions/load"
	"SmallVM/instructions/math"
	"SmallVM/instructions/store"
	"SmallVM/rtarea"
)

// WIDE is used to extend the range of local variables available to the instruction from 8 bits (i.e. 0-255) to 16 bits (i.e. 0-65535).
type WIDE struct {
	instruction common.Instruction
}

// For aload, dload, iload, fload, lload, astore, dstore, istore, fstore, lstore, and ret.
func (self *WIDE) FetchOperands(reader *common.ByteCodeReader) {
	bytecode := reader.ReadUInt8()
	switch bytecode {
	case 0x15: // iload
		inst := &load.ILOAD{}
		index := uint(reader.ReadUInt16())
		inst.SetIndex(index)
		self.instruction = inst
	case 0x16: // lload
		inst := &load.LLOAD{}
		index := uint(reader.ReadUInt16())
		inst.SetIndex(index)
		self.instruction = inst
	case 0x17: // fload
		inst := &load.FLOAD{}
		index := uint(reader.ReadUInt16())
		inst.SetIndex(index)
		self.instruction = inst
	case 0x18: // dload
		inst := &load.DLOAD{}
		index := uint(reader.ReadUInt16())
		inst.SetIndex(index)
		self.instruction = inst
	case 0x19: // aload
		inst := &load.ALOAD{}
		index := uint(reader.ReadUInt16())
		inst.SetIndex(index)
		self.instruction = inst
	case 0x36: // istore
		inst := &store.ISTORE{}
		index := uint(reader.ReadUInt16())
		inst.SetIndex(index)
		self.instruction = inst
	case 0x37: // lstore
		inst := &store.LSTORE{}
		index := uint(reader.ReadUInt16())
		inst.SetIndex(index)
		self.instruction = inst
	case 0x38: // fstore
		inst := &store.FSTORE{}
		index := uint(reader.ReadUInt16())
		inst.SetIndex(index)
		self.instruction = inst
	case 0x39: // dstore
		inst := &store.DSTORE{}
		index := uint(reader.ReadUInt16())
		inst.SetIndex(index)
		self.instruction = inst
	case 0x3a: // astore
		inst := &store.ASTORE{}
		index := uint(reader.ReadUInt16())
		inst.SetIndex(index)
		self.instruction = inst
	case 0x84: // iinc
		inst := &math.IINC{}
		index := uint(reader.ReadUInt16())
		increment := int32(reader.ReadInt16())
		inst.SetIndex(index)
		inst.SetIncrement(increment)
		self.instruction = inst
	case 0xa9: // ret
		panic("Unsupported opcode: 0xa9(ret)!")
	}
}

func (self *WIDE) Execute(frame *rtarea.Frame) {
	self.instruction.Execute(frame)
}
