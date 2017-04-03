package control

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

// branch to address
type GOTO struct {
	offset int16
}

func (self *GOTO) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt16()
}

func (self *GOTO) Execute(frame *rtarea.Frame) {
	pc := frame.CurrentThread().PC() - 3
	nextPC := pc + int(self.offset)
	frame.CurrentThread().SetPC(nextPC)
}
