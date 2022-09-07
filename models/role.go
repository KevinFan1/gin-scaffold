package models

type Role struct {
	BaseModel
	Name string `gorm:"type:varchar(32);not null;default:'';comment:角色名称" json:"name"`
	Code string `gorm:"type:varchar(32);not null;default:'';comment:角色code(CASBIN限制)'" json:"code"`
	//Permission string `gorm:"type:varchar(255);not null;default:'';comment:角色权限;" json:"permission"`
}

func (Role) TableName() string {
	return "t_role"
}
