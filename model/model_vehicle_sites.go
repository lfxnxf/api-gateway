package model

const (
	VehicleSitesTableName = "vehicle_sites"

	VehicleSitesStatusNormal  = 1 // 正常
	VehicleSitesStatusDeleted = 2 // 删除
)

type VehicleSitesModel struct {
	Id        int64  `json:"id" gorm:"column:id"`
	VehicleId int64  `json:"vehicle_id" gorm:"column:vehicle_id"`   // 车辆id
	SiteId    int64  `json:"site_id" gorm:"column:site_id"`         // 站点id
	Sort      int64  `json:"sort" gorm:"column:sort"`               // 排序
	Status    *int64 `json:"status" gorm:"column:status;default:1"` // 状态，1：正常，2：删除
}

func (v *VehicleSitesModel) TableName() string {
	return VehicleSitesTableName
}
