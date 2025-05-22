package constants

// 存放状态码 ：这个还是有好处的
const (
	// 1x : 服务器内部问题
	InternalError = 10
	// 2x : 外部问题
	FaultParams = 20

	// 3x : 成功信息
	Success = 30
)
