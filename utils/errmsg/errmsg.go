package errmsg

import "fmt"

/*
	处理程序中遇到的错误
*/

// 状态码
const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	// code = 2000... 文章模块错误
	ERROR_ART_NOT_EXIST = 2001

	// code = 3000... 分类模块错误
	ERROR_CATEGORY_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

// 根据不同的状态码，创建不同的信息
var codemsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误！",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_USER_NO_RIGHT:    "该用户无权限",

	ERROR_ART_NOT_EXIST: "文章不存在",

	ERROR_CATEGORY_USED:  "该分类已存在",
	ERROR_CATE_NOT_EXIST: "该分类不存在",
}

// 根据状态码，在codemsg中找到返回信息
func GetErrMsg(code int) string {
	msg, exist := codemsg[code]
	if exist {
		return msg
	}
	return fmt.Sprintf("不存在状态码: %d ", code)
}
