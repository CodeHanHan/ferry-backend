package db

import "errors"

var (
	ErrNotExist       error = errors.New("记录不存在")
	ErrDuplicateValue error = errors.New("记录值重复")
)
