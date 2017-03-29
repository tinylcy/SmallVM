package rtarea

import "math"

type LocalVariables struct {
	table []Slot
}

func NewLocalVariables(maxLocals uint) *LocalVariables {
	if maxLocals > 0 {
		localVars := &LocalVariables{}
		localVars.table = make([]Slot, maxLocals)
		return localVars
	}
	return nil
}

func (self *LocalVariables) SetInt(index uint, value int32) {
	self.table[index].Num = value
}

func (self *LocalVariables) GetInt(index uint) int32 {
	return self.table[index].Num
}

func (self *LocalVariables) SetFloat(index uint, value float32) {
	uint32Val := math.Float32bits(value)
	self.table[index].Num = int32(uint32Val)
}

func (self *LocalVariables) GetFloat(index uint) float32 {
	uint32Val := uint32(self.table[index].Num)
	return math.Float32frombits(uint32Val)
}

func (self *LocalVariables) SetLong(index uint, value int64) {
	low := int32(value)
	high := int32(value >> 32)
	self.table[index].Num = low
	self.table[index+1].Num = high
}

func (self *LocalVariables) GetLong(index uint) int64 {
	low := uint32(self.table[index].Num)
	high := uint32(self.table[index+1].Num)
	return int64(high)<<32 | int64(low)
}

func (self *LocalVariables) SetDouble(index uint, value float64) {
	int64Val := int64(math.Float64bits(value))
	self.SetLong(index, int64Val)
}

func (self *LocalVariables) GetDouble(index uint) float64 {
	uint64Val := uint64(self.GetLong(index))
	return math.Float64frombits(uint64Val)
}

func (self *LocalVariables) SetReference(index uint, ref *Object) {
	self.table[index].Ref = ref
}

func (self *LocalVariables) GetReference(index uint) *Object {
	return self.table[index].Ref
}
