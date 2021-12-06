package error_code

import (
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
)

//用户 错误码 全局102开头
var (
	HasEqPhone          = school_errors.AddError(102000001, "手机号重复！")
	HasNotAuthAddAdmin  = school_errors.AddError(102000002, "只有老板有权限添加管理员!")
	HasNotAuthAddDriver = school_errors.AddError(102000003, "只有老板和管理员有权限添加司机!")
	CannotEditPhone     = school_errors.AddError(102000004, "不可更改手机号!")
)
