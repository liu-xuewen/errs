package errx

import "github.com/liu-xuewen/errs"

// 101 err code

// DB DB
func DB() *errs.ErrWrap {
	return errs.New(101001, "数据操作失败")
}

// Create Create
func Create() *errs.ErrWrap {
	return errs.New(101002, "数据创建失败")
}

// Update Update
func Update() *errs.ErrWrap {
	return errs.New(101003, "数据更新失败")
}

// Delete Delete
func Delete() *errs.ErrWrap {
	return errs.New(101004, "数据删除失败")
}

// First First
func First() *errs.ErrWrap {
	return errs.New(101005, "数据查找失败")
}

// RecordNotFound RecordNotFound
func RecordNotFound() *errs.ErrWrap {
	return errs.New(101006, "数据不存在")
}
