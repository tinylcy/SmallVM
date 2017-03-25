// SmallVM project main.go
package main

import "SmallVM/cmd"
import "SmallVM/classpath"
import "SmallVM/classfile"
import "SmallVM/classfile/reader"
import "fmt"

func main() {
	cmd := cmd.ParseCmdLine()
	cp := classpath.Parse(cmd)
	data, _ := cp.ReadClass("java.lang.Object")
	fmt.Println(data)

	reader := reader.NewClassReader(data)
	cf := classfile.NewClassFile()
	cf.Parse(reader)
}
