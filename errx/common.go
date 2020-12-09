package errx

import "github.com/liu-xuewen/errs"

// Param Param
func Param() *errs.ErrWrap {
	return errs.New(100001, "参数异常")
}
