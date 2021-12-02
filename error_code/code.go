package error_code

import (
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
)

//common 错误码 全局100开头
var (
	UnIdentify = school_errors.AddError(100000001, "没有实名认证")
)

func Import() {

}
