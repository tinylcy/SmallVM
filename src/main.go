// SmallVM project main.go
package main

import (
	//	"SmallVM/classfile"
	//	"SmallVM/classfile/reader"
	//	"SmallVM/classpath"
	//	"SmallVM/cmd"
	"SmallVM/rtarea"
	"fmt"
)

func main() {
	//	cmd := cmd.ParseCmdLine()
	//	cp := classpath.Parse(cmd)
	//	data, _ := cp.ReadClass("java.lang.Double")
	//	reader := reader.NewClassReader(data)
	//	cf := classfile.NewClassFile()
	//	cf.Parse(reader)
	//	cf.ViewClass()

	// Test Local Variables
	localVars := rtarea.NewLocalVariables(10)
	localVars.SetInt(2, 24)
	fmt.Println(*localVars)
	localVars.SetLong(3, 22337203685477)
	fmt.Println(*localVars)
	fmt.Println(localVars.GetLong(3))
	localVars.SetDouble(6, 22337203685477.7)
	fmt.Println(*localVars)
	fmt.Println(localVars.GetDouble(6))

	// Test Operand Stack
	operandStack := rtarea.NewOperandStack(10)
	operandStack.PushLong(int64(22337203685477))
	fmt.Println(operandStack)
	fmt.Println(int64(operandStack.PopLong()))
}
