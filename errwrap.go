package errs

import (
	"runtime"
	"strconv"
)

var DefaultCode = 999999
var DefaultMsg = "操作失败"

// ErrWrap ErrWrap队列挂载err
type ErrWrap struct {
	Code   int           `json:"code"`
	Msg    string        `json:"msg"`
	Err    error         `json:"err"`
	Params []interface{} `json:"params"`
	Path   string        `json:"path"`
	Next   *ErrWrap      `json:"next"`
}

func (errWrap *ErrWrap) Error() string {
	return errWrap.Msg
}

// New New
func New(code int, msg string) *ErrWrap {
	e := new(ErrWrap)
	e.Code = code
	e.Msg = msg
	return e
}

func NewDefault() *ErrWrap {
	e := new(ErrWrap)
	e.Code = DefaultCode
	e.Msg = DefaultMsg
	return e
}

// Wrap Wrap
func Wrap(err error, params ...interface{}) *ErrWrap {
	errWrap := NewDefault()
	_, file, line, _ := runtime.Caller(1)
	errWrap.Path = file + strconv.Itoa(line)
	errWrap.Params = params

	er, ok := err.(*ErrWrap)
	if ok {
		er.Next = errWrap
		return er
	}

	errWrap.Err = err
	return errWrap
}

// Wrap Wrap
func (errWrap *ErrWrap) Wrap(err error, params ...interface{}) *ErrWrap {
	_, file, line, _ := runtime.Caller(1)
	errWrap.Path = file + strconv.Itoa(line)
	errWrap.Params = params

	er, ok := err.(*ErrWrap)
	if ok {
		er.Next = errWrap
		er.Err = err
		return er
	}
	errWrap.Err = err
	return errWrap
}

// Msg Msg
func Msg(msg string, params ...interface{}) *ErrWrap {
	errWrap := NewDefault()
	_, file, line, _ := runtime.Caller(1)
	errWrap.Path = file + strconv.Itoa(line)
	errWrap.Params = params
	errWrap.Msg = msg
	return errWrap
}

// Bind Bind
func (errWrap *ErrWrap) Bind(params ...interface{}) *ErrWrap {
	errWrap.Params = append(errWrap.Params, params)
	return errWrap
}

// Log Log
func (errWrap *ErrWrap) Log() ErrWrapLog {
	return ErrWrapLog(errWrap)
}
