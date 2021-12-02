package model

const (
	SitesTableName = "sites"

	SitesStatusNormal  = 1 // 正常
	SitesStatusDeleted = 2 // 删除
)

type SitesModel struct {
	Id        int64  `json:"id" gorm:"column:id"`
	Name      string `json:"name" gorm:"column:name"`               // 名称
	Longitude string `json:"longitude" gorm:"column:longitude"`     // 经度
	Latitude  string `json:"latitude" gorm:"column:latitude"`       // 纬度
	Status    *int64 `json:"status" gorm:"column:status;default:1"` // 状态，1：正常，2：删除
}

func (s *SitesModel) TableName() string {
	return SitesTableName
}
