// SmallVM project main.go
package main

import (
	"SmallVM/classfile"
	"SmallVM/classfile/reader"
	"SmallVM/classpath"
	"SmallVM/cmd"
)

func main() {
	cmd := cmd.ParseCmdLine()
	cp := classpath.Parse(cmd)
	data, _ := cp.ReadClass("java.lang.Double")
	reader := reader.NewClassReader(data)
	cf := classfile.NewClassFile()
	cf.Parse(reader)
	cf.ViewClass()
}
