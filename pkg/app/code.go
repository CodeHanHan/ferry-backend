package app

import "fmt"

type ErrCode int

var (
	// 客户端发送的数据包含非法参数
	Err_Invalid_Argument ErrCode = 400

	// token缺失或过期，请求未能通过身份认证
	Err_Unauthenticated ErrCode = 401

	// 客户端没有足够的权限
	Err_Permission_Denied ErrCode = 403

	// 资源未找到
	Err_Not_found ErrCode = 404

	// 并发冲突
	Err_Aborted ErrCode = 409

	// 资源配额不足或达不到速率限制
	Err_Resource_Exhausted ErrCode = 429

	// 请求被客户端取消
	Err_Cancelled ErrCode = 499

	// 内部服务器错误
	Err_Internal ErrCode = 500

	// API方法没有被服务器实现
	Err_Not_Implemented ErrCode = 501

	// 服务不可用
	Err_Unavailable ErrCode = 503

	// 请求超过了截止日期
	Err_Dealine_Exceed ErrCode = 504
)

func (c ErrCode) String() string {
	switch c {
	case Err_Invalid_Argument:
		return "Invalid Argument"
	case Err_Unauthenticated:
		return "Unauthenticated"
	case Err_Permission_Denied:
		return "Permission Denied"
	case Err_Not_found:
		return "Not found"
	case Err_Aborted:
		return "Aborted"
	case Err_Resource_Exhausted:
		return "Resource Exhausted"
	case Err_Internal:
		return "Internal Server Error"
	case Err_Not_Implemented:
		return "Not Implemented"
	case Err_Unavailable:
		return "service is unavailable"
	case Err_Dealine_Exceed:
		return "dealine exceeded"
	default:
		return "Unknown error"
	}
}

func (c ErrCode) Error() string {
	return fmt.Sprintf("%d: %s", c, c.String())
}
