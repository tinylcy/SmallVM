package common

import "SmallVM/rtarea"

type Instruction interface {
	FetchOperands(reader *ByteCodeReader)
	Execute(frame *rtarea.Frame)
}
