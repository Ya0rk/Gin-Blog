package errmsg

import "fmt"

/*
	处理程序中遇到的错误
*/

// 状态码
const (
	SUCC  = 200
	ERROR = 500

	// code = 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007

	// code = 2000... 文章模块错误
	// code = 3000... 分类模块错误
)

// 根据不同的状态码，创建不同的信息
var codemsg = map[int]string{
	SUCC:  "OK",
	ERROR: "FAIL",

	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误！",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
}

// 根据状态码，在codemsg中找到返回信息
func GetErrMeg(code int) string {
	msg, exist := codemsg[code]
	if exist {
		return msg
	}
	fmt.Println("不存在状态码: %d ", code)
}
