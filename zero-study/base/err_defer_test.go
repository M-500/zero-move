package base

import (
	"errors"
	"fmt"
	"testing"
)

// @Description
// @Author 代码小学生王木木

/*
如果一个函数的返回值是error，里面执行了多个defer，并且这些defer里面调用了不同的方法，也会返回error，
但是这些error的格式是不一样的（比如有一些方法返回的是官方的errors，有一些是业务定义的错误，比如错误码和错误信息）。
怎么样能统一处理这些defer的错误并且返回？

在Go语言中，处理多个defer中的错误并统一返回，可以通过一个常见的模式：将所有的错误收集起来，并在函数返回之前统一处理这些错误。以下是一个详细的步骤和示例，展示如何实现这种统一处理：
步骤
 1. 定义一个错误收集器：在函数开始时定义一个变量用于收集所有的错误。
 1. 使用defer来捕获错误：在每个可能产生错误的defer中，捕获并记录错误。
 1. 统一处理错误：在函数返回之前统一处理收集到的错误，并根据需要返回单一的错误信息。
*/
type CustomError struct {
	Code    int
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
func firstOperation() error {
	// Simulating an error
	return errors.New("第一步出错！firstOperation")
}

func secondOperation() error {
	// Simulating an error
	return CustomError{Code: 500, Message: "第二步出错！CustomError"}
}

func thirdOperation() error {
	// Simulating no error
	return nil
}
func combineErrors(errs []error) error {
	// You can customize this function to combine errors in a format you prefer
	var combinedErr string
	for _, e := range errs {
		combinedErr += e.Error() + "; "
	}
	return errors.New(combinedErr)
}

func doSomething() (err error) {
	var errCollector []error

	defer func() {
		if e := firstOperation(); e != nil {
			errCollector = append(errCollector, e)
		}
	}()

	defer func() {
		if e := secondOperation(); e != nil {
			errCollector = append(errCollector, e)
		}
	}()

	defer func() {
		if e := thirdOperation(); e != nil {
			errCollector = append(errCollector, e)
		}
	}()
	// Combine errors if any
	if len(errCollector) > 0 {
		err = combineErrors(errCollector)
		return err
	}
	return err
}

func TestErrorDemo(t *testing.T) {
	err := doSomething()
	if err != nil {
		fmt.Println("有一步出错了:", err)
	} else {
		fmt.Println("都成功了！")
	}
}
