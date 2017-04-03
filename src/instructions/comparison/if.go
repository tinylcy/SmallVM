package comparison

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// Jump if zero
type IFEQ struct {
	offset int16 // offset is a 16-bit signed integer
}

// Jump if nonzero
type IFNE struct {
	offset int16
}

// Jump if less than zero
type IFLT struct {
	offset int16
}

// Jump if greater than or equal to zero
type IFGE struct {
	offset int16
}

// Jump if greater than zero
type IFGT struct {
	offset int16
}

// Jump if less than or equal to zero
type IFLE struct {
	offset int16
}

func (self *IFEQ) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

// If the top value(int) equals zero, execution branches to the address (pc + offset).
func (self *IFEQ) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopInt()
	if value == 0 {
		pc := frame.CurrentThread().PC() - 3
		nextPC := pc + int(self.offset)
		frame.CurrentThread().SetPC(nextPC)
	}
}

func (self *IFNE) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

// If the top value(int) does not equal zero, execution branches to the address (pc + offset).
func (self *IFNE) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopInt()
	if value != 0 {
		pc := frame.CurrentThread().PC() - 3
		nextPC := pc + int(self.offset)
		frame.CurrentThread().SetPC(nextPC)
	}
}

func (self *IFLT) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

// If the top value(int) is less than zero, execution branches to the address (pc + offset).
func (self *IFLT) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopInt()
	if value < 0 {
		pc := frame.CurrentThread().PC() - 3
		nextPC := pc + int(self.offset)
		frame.CurrentThread().SetPC(nextPC)
	}
}

func (self *IFGE) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

// If the top value(int) is greater than or equal to zero, execution branches to the address (pc + offset).
func (self *IFGE) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopInt()
	if value >= 0 {
		pc := frame.CurrentThread().PC() - 3
		nextPC := pc + int(self.offset)
		frame.CurrentThread().SetPC(nextPC)
	}
}

func (self *IFGT) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

//  If the top value(int) is greater than zero, execution branches to the address (pc + offset).
func (self *IFGT) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopInt()
	if value > 0 {
		pc := frame.CurrentThread().PC() - 3
		nextPC := pc + int(self.offset)
		frame.CurrentThread().SetPC(nextPC)
	}
}

func (self *IFLE) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

// If the top value(int) is less than or equal to zero, execution branches to the address (pc + offset).
func (self *IFLE) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value := stack.PopInt()
	if value <= 0 {
		pc := frame.CurrentThread().PC() - 3
		nextPC := pc + int(self.offset)
		frame.CurrentThread().SetPC(nextPC)
	}
}
