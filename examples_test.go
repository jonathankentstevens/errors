package exception_test

import (
	"classes/Logger"
	"errors"
)

func ExampleHandleError() {
	err := errors.New("Test error")
	Logger.HandleError(Logger.ERROR_TYPE_CODE_ALERT, err)
}

func ExampleFileWrite() {
	err := errors.New("Test error")
	Logger.FileWrite(Logger.FILE_TEST, err)
}