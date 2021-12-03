package error_code

import (
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
)

//common 错误码 全局100开头
var (
	UnLogin               = school_errors.AddError(100000001, "请登录！")
	VerificationCodeWrong = school_errors.AddError(100000002, "请输入正确的验证码！")
	NotHasAuth            = school_errors.AddError(100000003, "没有操作权限！")
)

func Import() {

}
