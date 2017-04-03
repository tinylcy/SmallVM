package comparison

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
	"math"
)

// Long integer comparison
type LCMP struct {
}

// Single-precision float comparison (-1 on NaN)
type FCMPL struct {
}

// Single-precision float comparison (1 on NaN)
type FCMPG struct {
}

// Compare two doubles (-1 on NaN)
type DCMPL struct {
}

// Compare two doubles (1 on NaN)
type DCMPG struct {
}

func (self *LCMP) FetchOperands(reader *common.ByteCodeReader) {
}

// If the two long integers are the same, the int 0 is pushed onto the stack.
// If value2 is greater than value1, the int 1 is pushed onto the stack.
// If value1 is greater than value2, the int -1 is pushed onto the stack.
func (self *LCMP) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopLong()
	value2 := stack.PopLong()
	var result int32
	if value1 == value2 {
		result = 0
	} else if value1 < value2 {
		result = 1
	} else {
		result = -1
	}
	stack.PushInt(result)
}

func (self *FCMPL) FetchOperands(reader *common.ByteCodeReader) {
}

// If the two numbers are the same, the integer 0 is pushed onto the stack.
// If value2 is greater than value1, the integer 1 is pushed onto the stack.
// If value1 is greater than value2, the integer -1 is pushed onto the stack.
// If either number is NaN, the integer -1 is pushed onto the stack. +0.0 and -0.0 are treated as equal.
func (self *FCMPL) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopFloat()
	value2 := stack.PopFloat()
	var result int32
	if math.IsNaN(float64(value1)) || math.IsNaN(float64(value2)) {
		result = -1
	} else if value1 == value2 {
		result = 0
	} else if value1 < value2 {
		result = 1
	} else {
		result = -1
	}
	stack.PushInt(result)
}

func (self *FCMPG) FetchOperands(reader *common.ByteCodeReader) {
}

// If the two numbers are the same, the integer 0 is pushed onto the stack.
// If value2 is greater than value1, the integer 1 is pushed onto the stack.
// If value1 is greater than value2, the integer -1 is pushed onto the stack.
// If either number is NaN, the integer 1 is pushed onto the stack. +0.0 and -0.0 are treated as equal.
func (self *FCMPG) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopFloat()
	value2 := stack.PopFloat()
	var result int32
	if math.IsNaN(float64(value1)) || math.IsNaN(float64(value2)) {
		result = 1
	} else if value1 == value2 {
		result = 0
	} else if value1 < value2 {
		result = 1
	} else {
		result = -1
	}
	stack.PushInt(result)
}

func (self *DCMPL) FetchOperands(reader *common.ByteCodeReader) {
}

// If the two numbers are the same, the 32-bit integer 0 is pushed onto the stack.
// If value2 is greater than value1, the integer 1 is pushed onto the stack.
// If value1 is greater than value2, -1 is pushed onto the stack.
// If either numbers is NaN, the integer -1 is pushed onto the stack. +0.0 and -0.0 are treated as equal.
func (self *DCMPL) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopDouble()
	value2 := stack.PopDouble()
	var result int32
	if math.IsNaN(value1) || math.IsNaN(value2) {
		result = -1
	} else if value1 == value2 {
		result = 0
	} else if value1 < value2 {
		result = 1
	} else {
		result = -1
	}
	stack.PushInt(result)
}

func (self *DCMPG) FetchOperands(reader *common.ByteCodeReader) {
}

// If the two numbers are the same, the int 0 is pushed onto the stack.
// If value2 is greater than value1, the int 1 is pushed onto the stack.
// If value1 is greater than value2, -1 is pushed onto the stack.
// If either numbers is NaN, the int 1 is pushed onto the stack. +0.0 and -0.0 are treated as equal.
func (self *DCMPG) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	value1 := stack.PopDouble()
	value2 := stack.PopDouble()
	var result int32
	if math.IsNaN(value1) || math.IsNaN(value2) {
		result = 1
	} else if value1 == value2 {
		result = 0
	} else if value1 < value2 {
		result = 1
	} else {
		result = -1
	}
	stack.PushInt(result)
}
