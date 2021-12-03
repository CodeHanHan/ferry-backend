package db

import "errors"

var (
	ErrNotExist error = errors.New("记录不存在")
)
