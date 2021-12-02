package error_code

import (
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
)

//车辆 错误码 全局101开头
var (
	HasEqLicensePlate = school_errors.AddError(101000001, "车牌号重复！")
)

