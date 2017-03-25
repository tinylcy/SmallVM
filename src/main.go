// SmallVM project main.go
package main

import (
	"SmallVM/classpath"
	"SmallVM/cmd"
	"fmt"
)

func main() {
	cmd := cmd.ParseCmdLine()
	cp := classpath.Parse(cmd)
	data, _ := cp.ReadClass("com.sun.nio.zipfs.JarFileSystemProvider")
	fmt.Println(data)
}
