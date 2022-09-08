package models

type LoginRecord struct {
	ID       uint     `json:"id" gorm:"primarykey;comment:标识符"`
	UserId   uint     `json:"-" gorm:"column:user_id;type:int;comment:登录用户id"`
	User     *User    `gorm:"foreignKey:UserId;" json:"user,omitempty"`
	Agent    string   `json:"agent" gorm:"column:agent;type:varchar(255);comment:用户浏览器agent"`
	Ip       string   `json:"ip" gorm:"column:ip;type:varchar(16);comment:登录ip"`
	Status   string   `json:"status" gorm:"column:status;type:varchar(6);comment:http状态码"`
	Code     string   `json:"code" gorm:"column:code;type:varchar(6);comment:业务状态码"`
	Response string   `json:"response" gorm:"column:response;type:varchar(32);comment:响应内容"`
	LoginAt  JSONTime `json:"login_at" gorm:"column:login_at;comment:登录于"`
}

func (LoginRecord) TableName() string {
	return "t_login_record"
}

type OperationRecord struct {
	BaseModel
	UserId   uint   `json:"-" gorm:"column:user_id;type:int;comment:操作用户id"`
	Agent    string `json:"agent" gorm:"column:agent;type:varchar(255);comment:用户浏览器agent"`
	Ip       string `json:"ip" gorm:"column:ip;type:varchar(16);comment:操作ip"`
	Path     string `json:"path" gorm:"column:path;type:varchar(100);comment:路由地址"`
	Method   string `json:"method" gorm:"column:method;type:varchar(6);comment:操作方法"`
	Content  string `json:"content" gorm:"column:content;type:varchar(255);comment:提交内容"`
	Status   string `json:"status" gorm:"column:status;type:varchar(6);comment:http状态码"`
	Code     string `json:"code" gorm:"column:code;type:varchar(6);comment:业务状态码"`
	Response string `json:"response" gorm:"column:response;type:varchar(32);comment:响应内容"`
	Cost     string `json:"cost" gorm:"column:cost;type:varchar(6);comment:总共耗费时间;not null;default:0"`
}

func (OperationRecord) TableName() string {
	return "t_operation_record"
}
