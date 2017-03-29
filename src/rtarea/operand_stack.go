package rtarea

import "math"

type OperandStack struct {
	top   uint
	slots []Slot
}

func NewOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		stack := &OperandStack{}
		stack.top = 0
		stack.slots = make([]Slot, maxStack)
		return stack
	}
	return nil
}

func (self *OperandStack) PushInt(value int32) {
	self.slots[self.top].Num = value
	self.top++
}

func (self *OperandStack) PopInt() int32 {
	self.top--
	value := self.slots[self.top].Num
	return value
}

func (self *OperandStack) PushFloat(value float32) {
	int32Val := int32(math.Float32bits(value))
	self.slots[self.top].Num = int32Val
	self.top++
}

func (self *OperandStack) PopFloat() float32 {
	self.top--
	uint32Val := uint32(self.slots[self.top].Num)
	return math.Float32frombits(uint32Val)
}

func (self *OperandStack) PushLong(value int64) {
	low := int32(value)
	high := int32(value >> 32)
	self.slots[self.top].Num = low
	self.slots[self.top+1].Num = high
	self.top += 2
}

func (self *OperandStack) PopLong() int64 {
	self.top--
	high := uint32(self.slots[self.top].Num)
	self.top--
	low := uint32(self.slots[self.top].Num)
	return int64(high)<<32 | int64(low)
}

func (self *OperandStack) PushDouble(value float64) {
	int64Val := int64(math.Float64bits(value))
	self.PushLong(int64Val)
}

func (self *OperandStack) PopDouble() float64 {
	uint64Val := uint64(self.PopLong())
	return math.Float64frombits(uint64Val)
}

func (self *OperandStack) PushReference(value *Object) {
	self.slots[self.top].Ref = value
	self.top++
}

func (self *OperandStack) PopReference() *Object {
	self.top--
	value := self.slots[self.top].Ref
	self.slots[self.top].Ref = nil
	return value
}
