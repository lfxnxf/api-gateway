package model

const (
	ShiftsTableName = "shifts"

	ShiftsStatusNormal  = 1 // 正常
	ShiftsStatusDeleted = 2 // 删除
)

type ShiftsModel struct {
	Id     int64  `json:"id" gorm:"column:id"`
	Name   string `json:"name" gorm:"column:name"`
	Status *int64 `json:"status" gorm:"column:status;default:1"` // 状态，1：正常，2：删除
}

func (s *ShiftsModel) TableName() string {
	return ShiftsTableName
}
