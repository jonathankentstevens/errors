[![License](http://img.shields.io/:license-gpl3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0.html)
[![Go Report Card](https://goreportcard.com/badge/github.com/jrkt/errors)](https://goreportcard.com/report/github.com/jrkt/errors)
[![GoDoc](https://godoc.org/github.com/jrkt/errors?status.svg)](https://godoc.org/github.com/jrkt/errors)
[![Build Status](https://travis-ci.org/jrkt/errors.svg?branch=master)](https://travis-ci.org/jrkt/errors)

# errors

Simple error handler implementation for GoLang while maintaining the idiomatic standard

# implementation
	go get github.com/jrkt/errors
	
The new error package is called just the same as the standard library error package:
          
	errors.New("test error")
	
	
# usage
```go
package main

import (
	  "github.com/jrkt/errors"
)

func main() {
    err := testFunction()
    if err != nil {
        println(err.Error()) 
    }
    
    err = errors.Errorf("new %s error", 12345)
}

func testFunction() error {
	  return errors.New("test error")
}
```

# example

Instead of the usual error message string, you will get this from the example error from above:

```
main.testFunction [test.go:15] - test error
```
