package error_handling

import (
	"errors"
	"fmt"
)

type recoverFunc func()

func StandartErrorHandling() (err error) {
	err = errFunc1()
	if err != nil {
		return
	}

	err = errFunc2()
	if err != nil {
		return
	}

	return
}

func ErrorHandlingWithPanic() (err error) {
	defer func() {
		if e := recover(); e != nil {
			if _, ok := e.(error); !ok {
				msg := fmt.Sprintf("Something went wrong: %v", e)
				err = errors.New(msg)
			} else {
				err = e.(error)
			}
		}
	}()

	err = errFunc1()
	if err != nil {
		panic(err)
	}

	err = errFunc2()
	if err != nil {
		panic(err)
	}

	return
}

func ErrorHandlingWithPanicIfHelper() (err error) {
	defer func() {
		if e := recover(); e != nil {
			if _, ok := e.(error); !ok {
				msg := fmt.Sprintf("Something went wrong: %v", e)
				err = errors.New(msg)
			} else {
				err = e.(error)
			}
		}
	}()

	err = errFunc1()
	panicIf(err)

	err = errFunc2()
	panicIf(err)

	return
}

func ErrorHandlingWithPanicIfHelperAndRescueFactory() (err error) {
	defer rescue(err)()

	err = errFunc1()
	panicIf(err)

	err = errFunc2()
	panicIf(err)

	return
}

func ErrorHandlingWithRunner() (err error) {
	run, runner := newRunner(err)
	run(errFunc1)
	run(errFunc2)

	return runner.E
}

type runner struct {
	F func() error
	E error
}

func (r *runner) Run() {
	if r.E != nil {
		return
	}

	r.E = r.F()
}

var rr *runner

func newRunner(e error) (func(func() error), *runner) {
	rr = &runner{E: e}
	return try, rr
}

func try(f func() error) {
	rr.F = f
	rr.Run()
}

func runErr() error {
	return rr.E
}

func panicIf(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func rescue(err error) recoverFunc {
	return func() {
		if e := recover(); e != nil {
			if _, ok := e.(error); !ok {
				msg := fmt.Sprintf("Something went wrong: %v", e)
				err = errors.New(msg)
			} else {
				err = e.(error)
			}
		}
	}
}

func errFunc1() error {
	return errors.New("error1")
}

func errFunc2() error {
	return errors.New("error2")
}
