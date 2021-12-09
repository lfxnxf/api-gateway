package model

const (
	ShiftsSitesTableName = "shifts_sites"

	ShiftsSitesStatusNormal  = 1 // 正常
	ShiftsSitesStatusDeleted = 2 // 删除
)

type ShiftsSitesModel struct {
	Id         int64  `json:"id" gorm:"column:id"`
	ShiftId    int64  `json:"shift_id" gorm:"column:shift_id"`       // 班次id
	SiteId     int64  `json:"site_id" gorm:"column:site_id"`         // 站点id
	ArriveTime string `json:"arrive_time" gorm:"column:arrive_time"` // 到站时间
	Sort       int64  `json:"sort" gorm:"column:sort"`               // 排序
	Status     *int64 `json:"status" gorm:"column:status;default:1"` // 状态，1：正常，2：删除
}

func (s *ShiftsSitesModel) TableName() string {
	return ShiftsSitesTableName
}
