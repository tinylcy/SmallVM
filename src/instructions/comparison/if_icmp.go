package comparison

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Jump if two integers are equal
type IF_ICMPEQ struct {
	offset int16
}

// Jump if two integers are not equal
type IF_ICMPNE struct {
	offset int16
}

// Jump if one integer is less than another
type IF_ICMPLT struct {
	offset int16
}

// Jump if one integer is greater than or equal to another
type IF_ICMPGE struct {
	offset int16
}

// Jump if one integer is greater than another
type IF_ICMPGT struct {
	offset int16
}

// Jump if one integer is less than or equal to another
type IF_ICMPLE struct {
	offset int16
}

func (self *IF_ICMPEQ) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

func (self *IF_ICMPEQ) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopInt()
	value2 := stack.PopInt()
	if value1 == value2 {
		pc := frame.CurrentThread().PC() - 3
		nextPC := pc + int(self.offset)
		frame.CurrentThread().SetPC(nextPC)
	}
}

func (self *IF_ICMPNE) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

func (self *IF_ICMPNE) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopInt()
	value2 := stack.PopInt()
	if value1 != value2 {
		pc := frame.CurrentThread().PC() - 3
		nextPC := pc + int(self.offset)
		frame.CurrentThread().SetPC(nextPC)
	}
}

func (self *IF_ICMPLT) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

//  If value2 is less than value1, execution branches to the address (pc + offset).
func (self *IF_ICMPLT) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopInt()
	value2 := stack.PopInt()
	if value1 > value2 {
		pc := frame.CurrentThread().PC() - 3
		nextPC := pc + int(self.offset)
		frame.CurrentThread().SetPC(nextPC)
	}
}

func (self *IF_ICMPGE) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

//  If value2 is greater than or equal to value1, execution branches to the address (pc + offset).
func (self *IF_ICMPGE) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopInt()
	value2 := stack.PopInt()
	if value1 <= value2 {
		pc := frame.CurrentThread().PC() - 3
		nextPC := pc + int(self.offset)
		frame.CurrentThread().SetPC(nextPC)
	}
}

func (self *IF_ICMPGT) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

// If value2 is greater than value1, execution branches to the address (pc + offset).
func (self *IF_ICMPGT) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopInt()
	value2 := stack.PopInt()
	if value1 < value2 {
		pc := frame.CurrentThread().PC() - 3
		nextPC := pc + int(self.offset)
		frame.CurrentThread().SetPC(nextPC)
	}
}

func (self *IF_ICMPLE) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

//  If value2 is less than or equal to value1, execution branches to the address (pc + offset).
func (self *IF_ICMPLE) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopInt()
	value2 := stack.PopInt()
	if value1 >= value2 {
		pc := frame.CurrentThread().PC() - 3
		nextPC := pc + int(self.offset)
		frame.CurrentThread().SetPC(nextPC)
	}
}
