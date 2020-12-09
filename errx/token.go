package errx

import "github.com/liu-xuewen/errs"

// TokenInvalid TokenInvalid
func TokenInvalid() *errs.ErrWrap {
	return errs.New(401401, "token无效")
}

