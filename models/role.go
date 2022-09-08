package models

type Role struct {
	BaseModel
	Name      string `json:"name" gorm:"type:varchar(32);not null;default:'';comment:角色名称"`
	Code      string `json:"code" gorm:"type:varchar(32);not null;default:'';comment:角色code(CASBIN限制)"`
	Remark    string `json:"remark" gorm:"column:remark;type:varchar(255);not null;default='';comment:备注信息"`
	IsDisable bool   `json:"is_disable" gorm:"column:is_disable;type:tinyint;comment:是否禁用;default:0;"`
	IsDel     bool   `json:"is_del" gorm:"column:is_del;type:tinyint;comment:是否删除;default:0;"`
	//Permission string `gorm:"type:varchar(255);not null;default:'';comment:角色权限;" json:"permission"`
}

func (Role) TableName() string {
	return "t_role"
}
