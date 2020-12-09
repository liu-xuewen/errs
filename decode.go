package errs

import (
	"context"
	"runtime"
	"strconv"

	"github.com/liu-xuewen/logger"
)

var (
	Success = NewErrWrap(0, "操作成功")
	Fail    = NewErrWrap(999999, "操作失败")
)

type ErrWrapLog *ErrWrap

// DecodeErrWrap DecodeErrWrap
func DecodeErrWrap(ctx context.Context, err error) (int, string) {
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