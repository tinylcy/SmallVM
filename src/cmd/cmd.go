package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type CmdLine struct {
	help      bool
	version   bool
	jre       string
	classpath string
	class     string
	args      []string
}

func ParseCmdLine() *CmdLine {
	cmd := &CmdLine{}

	flag.Usage = cmd.PrintHelpMessage
	flag.BoolVar(&cmd.help, "help", false, "print help message")
	flag.BoolVar(&cmd.help, "?", false, "print help message")
	flag.BoolVar(&cmd.version, "version", false, "print version message")
	flag.StringVar(&cmd.jre, "jre", "", "bootstrap classpath and extension classpath")
	flag.StringVar(&cmd.classpath, "classpath", "", "classpath")
	flag.StringVar(&cmd.classpath, "cp", "", "classpath")

	flag.Parse()

	// parse class and args
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func (self *CmdLine) PrintHelpMessage() {
	fmt.Println("Usage: java [-options] class [args...]\n")
	fmt.Println("-version   输出产品版本并退出")
	fmt.Println("-? -help   输出此帮助消息")
	fmt.Println("-cp        <目录和 zip/jar 文件的类搜索路径>")
	fmt.Println("-classpath <目录和 zip/jar 文件的类搜索路径>")
}

func (self *CmdLine) PrintVersionMessage() {
	fmt.Println("java version \"0.0.1\"")
}

func (self *CmdLine) StartVM() {
	if self.version {
		self.PrintVersionMessage()
	} else if self.help || self.class == "" {
		self.PrintHelpMessage()
	} else {
		fmt.Printf("classpath: %s class:%s args: %v\n", self.classpath, self.class, self.args)
	}
}

// return JRE location(include bootstrap classpath and extension classpath)
func (self *CmdLine) JREPath() (string, error) {
	jrePath := self.jre
	if jrePath != "" && exists(jrePath) {
		return jrePath, nil
	}
	jh := os.Getenv("JAVA_HOME")
	if jh != "" {
		jh = filepath.Join(jh, "jre")
		return jh, nil
	}
	return "", errors.New("The system cannot find JRE path!")
}

// return user classpath, the default location is current path
func (self *CmdLine) AppClassPath() (string, error) {
	cp := self.classpath
	if cp != "" {
		return cp, nil
	}
	return ".", nil
}

// check whether the given file or directory exists or not
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
