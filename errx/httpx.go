package errx

import "github.com/liu-xuewen/errs"

// Http Http
func Http() *errs.ErrWrap {
	return errs.New(103001, "http请求异常")
}

