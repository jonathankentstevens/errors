//Package Logger is the main package for logging errors for all processes
package exception

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"time"
)

type ErrorDetail struct {
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

func (err *ErrorDetail) Error() string {
	return fmt.Sprintf("%s - %s.%s(%d): %s", err.File, err.Package, err.Method, err.Line, err.Msg)
}

func New(s string) error {
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

	return &ErrorDetail{Msg: s, File: file, Package: pkg, Method: method, Line: line, Ip: ip, Host: hostname, Date: time.Now()}
}
