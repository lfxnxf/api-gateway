// Generated by the inits tool.  DO NOT EDIT!
package http

import (
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
)

func initRoute(s httpserver.Server) {
	s.ANY("/ping", ping)

	s.GET("/ws", wsHandler) // 长连接服务

	s.POST("/api/login", login)                                       // 登录
	s.POST("/api/login/send_verification_code", sendVerificationCode) // 发送验证码

	s.GET("/api/user/get", getUserInfo)    // 获取用户资料
	s.POST("/api/user/edit", editUserInfo) // 修改用户资料

	s.POST("/api/v1/vehicle/save", saveVehicle)               // 新增车辆信息
	s.GET("/api/v1/vehicle_info/get", getAllVehicleInfo)      // 获取车辆分类
	s.POST("/api/v1/driver/add", addDriver)                   // 新增司机
	s.GET("/api/v1/driver/get", getDrivers)                   // 获取全部司机
	s.GET("/api/v1/identity/get_default", getDefaultIdentity) // 获取全部默认身份

	s.POST("/api/v1/vehicle/sites/save", saveVehicleSites) // 车辆保存站点

	s.GET("/api/v1/sites/get_list", getSitesList) // 获取站点

	s.GET("/api/v1/shifts/default/get", getDefaultShifts) // 获取默认班次

}
