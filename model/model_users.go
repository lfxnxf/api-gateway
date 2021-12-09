package model

const (
	UsersTableName = "users"

	IdentityBoss    = 1 // 主管理员
	IdentityAdmin   = 2 // 管理员
	IdentityDriver  = 3 // 司机
	IdentityParents = 4 // 家长

	UserStatusNormal  = 1 // 正常
	UserStatusDeleted = 2 // 删除
)

var identityTextMap = map[int64]string{
	IdentityBoss:    "主管理员",
	IdentityAdmin:   "管理员",
	IdentityDriver:  "乘务员",
	IdentityParents: "家长",
}

func GetIdentityText(identity int64) string {
	return identityTextMap[identity]
}

type UsersModel struct {
	Id       int64  `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`         // 用户名字
	Address  string `json:"address" gorm:"column:address"`   // 地址
	BossId   int64  `json:"boss_id" gorm:"column:boss_id"`   // 老板uid
	Identity int64  `json:"identity" gorm:"column:identity"` // 身份，1：老板，2：家长，3：司机
	Phone    int64  `json:"phone" gorm:"column:phone"`       // 手机号
	Password string `json:"password" gorm:"column:password"`
	Token    string `json:"token" gorm:"column:token"`
	Status   *int64 `json:"status" gorm:"column:status;default:1"` // 用户状态，1：正常，2：删除
}

func (u *UsersModel) TableName() string {
	return UsersTableName
}
