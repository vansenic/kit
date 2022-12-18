package code

type Code int

const CodeDefault = Success

const (
	Success    Code = 200 // 操作成功
	Check      Code = 300 // 请检查
	Timeout    Code = 301 // 超时
	ExpStatus  Code = 401 // 状态码
	ExpMarshal Code = 402 // 编码解码
	ExpRuntime Code = 403 // 运行时
	NotExist   Code = 404 // 资源不存在
	Unverified Code = 405 // 未通过验证
	ExpLimit   Code = 406 // 超限
	ErrRemote  Code = 500 // 远端错误
	ErrCall    Code = 501 // 调用错误
	Interrupt  Code = 502 // 中断
	Ignore     Code = 600 // 可忽略
	UseLess    Code = 700 // 无用数据
	Intercept  Code = 701 // 被拦截
)

var commentaries = map[Code]string{
	Success:    "操作成功",
	Check:      "请检查",
	Timeout:    "超时",
	ExpStatus:  "非常规状态码",
	ExpMarshal: "编码或映射错误",
	ExpRuntime: "运行时发生异常",
	NotExist:   "资源不存在",
	Unverified: "未通过验证",
	ExpLimit:   "超限",
	ErrRemote:  "远端错误",
	ErrCall:    "调用错误",
	Interrupt:  "中断",
	Ignore:     "可忽略",
	UseLess:    "无用数据",
	Intercept:  "被拦截",
}

func GetCommentaries(code Code) string {
	message := commentaries[code]
	if message == "" {
		return "未知错误"
	}
	return message
}
