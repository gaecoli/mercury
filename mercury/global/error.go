package global

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrorMix struct {
	ValidateArgsError       CustomError
	AuthorizationError      CustomError
	AccessDeniedError       CustomError
	NotFoundError           CustomError
	RequestTimeoutError     CustomError
	ServerInternalError     CustomError
	NetworkGatewayError     CustomError
	NotImplementedError     CustomError
	ServerNotAvailableError CustomError
	HttpVersionError        CustomError
}

var Errors = CustomErrorMix{
	ValidateArgsError:       CustomError{40000, "请求错误"},
	AuthorizationError:      CustomError{40100, "未授权，请登录"},
	AccessDeniedError:       CustomError{40300, "拒绝访问"},
	NotFoundError:           CustomError{40400, "未找到资源"},
	RequestTimeoutError:     CustomError{40800, "请求超时"},
	ServerInternalError:     CustomError{50000, "服务器内部错误"},
	NetworkGatewayError:     CustomError{50200, "网关错误"},
	NotImplementedError:     CustomError{50100, "服务未实现"},
	ServerNotAvailableError: CustomError{50300, "服务不可用"},
	HttpVersionError:        CustomError{505, "HTTP版本不支持"},
}
