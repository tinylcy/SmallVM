package stack

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type DUP struct {
}

type DUP_X1 struct {
}

type DUP_X2 struct {
}

type DUP2 struct {
}

type DUP2_X1 struct {
}

type DUP2_X2 struct {
}

func (self *DUP) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *DUP) Execute(frame *rtarea.Frame) {
	slot := frame.OperandStack().PeekSlot()
	slotCopy := slot
	frame.OperandStack().PushSlot(slotCopy)
}

func (self *DUP_X1) FetchOperands(reader *common.ByteCodeReader) {
}

// Duplicate top stack slot and insert beneath second slot.
func (self *DUP_X1) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (self *DUP_X2) FetchOperands(reader *common.ByteCodeReader) {
}

// Duplicate top stack slot and insert beneath third slot.
func (self *DUP_X2) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (self *DUP2) FetchOperands(reader *common.ByteCodeReader) {
}

// Duplicate top two stack slots
func (self *DUP2) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (self *DUP2_X1) FetchOperands(reader *common.ByteCodeReader) {
}

// Duplicate two slots and insert beneath third slot
func (self *DUP2_X1) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (self *DUP2_X2) FetchOperands(reader *common.ByteCodeReader) {
}

// Duplicate two slots and insert beneath fourth slot
func (self *DUP2_X2) Execute(frame *rtarea.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
