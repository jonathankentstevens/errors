[![License](http://img.shields.io/:license-gpl3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0.html)
[![Go Report Card](https://goreportcard.com/badge/github.com/jonathankentstevens/exception)](https://goreportcard.com/report/github.com/jonathankentstevens/exception)
[![GoDoc](https://godoc.org/github.com/jonathankentstevens/exception?status.svg)](https://godoc.org/github.com/jonathankentstevens/exception)
[![Build Status](https://travis-ci.org/jonathankentstevens/exception.svg?branch=master)](https://travis-ci.org/jonathankentstevens/exception)

# exception
Custom error handler for Golang while keeping the standard for error handling

# implementation
    go get github.com/jonathankentstevens/errors
The new error package is called just the same as the standard library error package:
          
     errors.New("test error")
    
# usage
```go
package main

import (
	  "github.com/jonathankentstevens/errors"
)

func main() {
    err := testFunction()
    if err != nil {
        println(err.Error()) 
    }
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
