package model

const (
	VehicleTableName     = "vehicle"
	VehicleInfoTableName = "vehicle_info"

	VehicleNormal  = 1 // 正常
	VehicleDeleted = 2 // 删除
)

type VehicleModel struct {
	Id            int64  `json:"id" gorm:"column:id" json:"id"`
	BossId        int64  `json:"boss_id" gorm:"column:boss_id" json:"boss_id"`                          // 车主id
	VehicleInfoId int64  `json:"vehicle_info_id"  gorm:"column:vehicle_info_id" json:"vehicle_info_id"` // 车辆类型Id
	LicensePlate  string `json:"license_plate"  gorm:"column:license_plate" json:"license_plate"`       // 车牌号
	DriverId      int64  `json:"driver_id"  gorm:"column:driver_id" json:"driver_id"`                   // 司机id
	Status        *int64 `json:"status" gorm:"column:status;default:1"`                                 // 状态，1：正常，2：删除
}

func (v *VehicleModel) TableName() string {
	return VehicleTableName
}

type VehicleInfoModel struct {
	Id      int64  `json:"id" gorm:"column:id" json:"id"`
	Name    string `json:"name" gorm:"column:name" json:"name"`              // 名称
	LoadNum int64  `json:"load_num"  gorm:"column:load_num" json:"load_num"` // 荷载人数
}

func (v *VehicleInfoModel) TableName() string {
	return VehicleInfoTableName
}
