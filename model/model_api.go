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

type GetSitesListReq struct {
	Name string `json:"name"`
}

type SiteInfo struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type GetSitesListResp struct {
	List []SiteInfo `json:"list"`
}

type SaveVehicleSitesReq struct {
	VehicleId int64             `json:"vehicle_id"`
	List      []VehicleSiteInfo `json:"list"`
}

type VehicleSiteInfo struct {
	VehicleId int64 `json:"vehicle_id"`
	SiteId    int64 `json:"site_id"`
	Sort      int64 `json:"sort"`
}

type EditUserInfoReq struct {
	Name string `json:"name"`
	Phone int64 `json:"phone"`
	Address string `json:"address"`
}
