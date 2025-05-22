package constants

// 存放状态码 ：这个还是有好处的
const (
	// 0x : 服务器内部问题
	InternalError = 01
	// 1x : 外部问题
	FaultParams = 10

	// 2x : 成功信息
	Success = 20
)
