# exception
Custom error handler for Golang while keeping the standard for error handling

# implementation
    go get github.com/jonathankentstevens/exception
    
The new exception package:
          
     exception.New("test error")
    
replaces the normal error package for initializing new error types:
     
     errors.New("test error")
    
# usage
```go
package main

import (
	  "exception"
)

func main() {
    err := testFunction()
    if err != nil {
        println(err.Error())
        //prints: ~/GO/src/test/test.go - main.testFunction(15): test error
    }
}

func testFunction() error {
	  return exception.New("test error")
}
```
