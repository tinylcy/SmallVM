package constant

import (
	"SmallVM/instructions/common"
	"SmallVM/rtarea"
)

type NOP struct {
}

func (self *NOP) FetchOperands(reader *common.ByteCodeReader) {
}

func (self *NOP) Execute(frame *rtarea.Frame) {
}
