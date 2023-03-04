package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	ErrorExistUser      = 102
	ErrorNotExistUser   = 103
	ErrorFailEncryption = 106
	ErrorNotCompare     = 107

	ErrorAuthCheckTokenFail    = 301 //token 错误
	ErrorAuthCheckTokenTimeout = 302 //token 过期
	ErrorAuthToken             = 303
	ErrorAuth                  = 304
	ErrorDatabase              = 401
	ErrorNotExistData          = 402
	ErrorNotAdmin              = 403
	ErrorExistInteraction      = 404
	ErrorNotExistInteraction   = 405
)

var MsgFlags = map[int]string{
	SUCCESS:       "成功",
	ERROR:         "失败!!!!",
	InvalidParams: "请求参数错误",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
	ErrorNotExistUser:          "用户不存在，请先注册",
	ErrorNotCompare:            "密码不匹配",
	ErrorDatabase:              "数据库操作出错,请重试",
	ErrorExistUser:             "用户已存在",
	ErrorFailEncryption:        "加密密码失败",
	ErrorNotExistData:          "数据不存在",

	ErrorNotAdmin:            "不是管理员",
	ErrorExistInteraction:    "该互动已存在",
	ErrorNotExistInteraction: "该互动不存在",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
