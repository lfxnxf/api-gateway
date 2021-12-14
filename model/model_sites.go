package model

const (
	SitesTableName = "sites"

	SitesStatusNormal  = 1 // 正常
	SitesStatusDeleted = 2 // 删除

	TypeNormal = 1 // 普通站点
	TypeSchool = 2 // 学校站点
)

type SitesModel struct {
	Id        int64  `json:"id" gorm:"column:id"`
	Name      string `json:"name" gorm:"column:name"`               // 名称
	Longitude string `json:"longitude" gorm:"column:longitude"`     // 经度
	Latitude  string `json:"latitude" gorm:"column:latitude"`       // 纬度
	Status    *int64 `json:"status" gorm:"column:status;default:1"` // 状态，1：正常，2：删除
	Type      int64  `json:"type" gorm:"column:type"`               // 类型，1：普通站点，2：学校站点
}

func (s *SitesModel) TableName() string {
	return SitesTableName
}
