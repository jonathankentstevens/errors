//Package exception is a custom error handler to provide better details about an error
package exception

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"time"
)

//Exception contains the important information pertaining to an error thrown
type Exception struct {
	Msg     string
	File    string
	Package string
	Method  string
	Line    int
	Ip      string
	Host    string
	Date    time.Time
	Stack   string
}

//New returns a new instance of Exception
func New(s string) Exception {
	hostname, _ := os.Hostname()
	ipArr, _ := net.LookupHost(hostname)
	var ip string
	if len(ipArr) == 1 {
		ip = ipArr[0]
	}
	pc, file, line, _ := runtime.Caller(1)
	path := runtime.FuncForPC(pc).Name()
	pathArgs := strings.Split(path, ".")
	pkg := pathArgs[0]
	method := pathArgs[1]

	return &Exception{Msg: s, File: file, Package: pkg, Method: method, Line: line, Ip: ip, Host: hostname, Date: time.Now()}
}


//Error returns a formatted string displaying the file where the error was thrown, the package it was
//thrown in, the method, line and, of course, the error message
func (err Exception) Error() string {
	return fmt.Sprintf("%s - %s.%s(%d): %s", err.File, err.Package, err.Method, err.Line, err.Msg)
}

//Stacktrace prints out the stacktrace of the error thrown
func (err Exception) Stacktrace() string {
	currentPkg := ""
	i := 1
	for {
		pc, file, line, _ := runtime.Caller(i)
		path := runtime.FuncForPC(pc).Name()
		pathArgs := strings.Split(path, ".")
		pkg := pathArgs[0]
		method := pathArgs[1]

		if pkg == "runtime" {
			break
		}

		if pkg != currentPkg {
			err.Stack += pkg
			currentPkg = pkg
		}

		err.Stack += fmt.Sprintf("\n\t%s - %s(%d)", file, method, line)
		err.Stack += "\n"

		i++
	}

	return err.Stack
}