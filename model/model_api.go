package model

type LoginReq struct {
	Phone            int64  `json:"phone"`
	VerificationCode string `json:"verification_code"`
}

type LoginResp struct {
	Token string          `json:"token"`
	User  GetUserInfoResp `json:"user"`
}

type SendVerificationCodeReq struct {
	Phone int64 `json:"phone"`
}

type AddDriverReq struct {
	Name     string `json:"name"`
	Phone    int64  `json:"phone"`
	Identity int64  `json:"identity"`
}

type AddVehicleReq struct {
	VehicleInfoId int64  `json:"vehicle_info_id"` // 车辆类型Id
	LicensePlate  string `json:"license_plate"`   // 车牌号
	DriverId      int64  `json:"driver_id"`       // 司机id
}

type VehicleInfo struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	LoadNum int64  `json:"load_num"`
}

type GetVehicleInfoResp struct {
	List []VehicleInfo `json:"list"`
}

type DriverInfo struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Phone    int64  `json:"phone"`
	Identity int64  `json:"identity"`
}

type GetDriversResp struct {
	List []DriverInfo `json:"list"`
}

type GetUserInfoResp struct {
	Uid      int64  `json:"uid"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    int64  `json:"phone"`
	Identity int64  `json:"identity"`
}
