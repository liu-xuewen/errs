package errs

import (
	"context"
	"runtime"
	"strconv"

	"github.com/liu-xuewen/logger"
)

var (
	Success = New(0, "操作成功")
	Fail    = New(999999, "操作失败")
)

type ErrWrapLog *ErrWrap

// DeErrCodeLog DeErrCodeLog
func DeErrCodeLog(ctx context.Context, err error) (int, string) {
	if err == nil {
		return Success.Code, Success.Msg
	}
	switch typed := err.(type) {
	case *ErrWrap:
		errWrap := err.(*ErrWrap)
		logger.Error(ctx, errWrap.Msg, "err_wrap", ErrWrapLog(errWrap))
		return typed.Code, typed.Msg
	default:
		_, file, line, _ := runtime.Caller(1)
		path := file + strconv.Itoa(line)
		logger.Error(ctx, "default", "err", err, "path", path)
	}

	return Fail.Code, Fail.Msg
}

// DeErrLog DeErrLog
func DeErrLog(ctx context.Context, err error) string {
	if err == nil {
		return ""
	}
	switch typed := err.(type) {
	case *ErrWrap:
		errWrap := err.(*ErrWrap)
		logger.Error(ctx, errWrap.Msg, "err_wrap", ErrWrapLog(errWrap))
		return typed.Msg + strconv.Itoa(typed.Code)
	default:
		_, file, line, _ := runtime.Caller(1)
		path := file + strconv.Itoa(line)
		logger.Error(ctx, "default", "err", err, "path", path)
	}

	return Fail.Msg + strconv.Itoa(Fail.Code)
}
