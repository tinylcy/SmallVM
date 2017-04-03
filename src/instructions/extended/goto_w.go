package extended

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type GOTO_W struct {
	offset int32 // offset is a 32-bit signed integer parameter
}

func (self *GOTO_W) FetchOperands(reader *common.ByteCodeReader) {
	self.offset = reader.ReadInt32()
}

func (self *GOTO_W) Execute(frame *rtarea.Frame) {
	pc := frame.CurrentThread().PC() - 5
	nextPC := pc + int(self.offset)
	frame.CurrentThread().SetPC(nextPC)
}
