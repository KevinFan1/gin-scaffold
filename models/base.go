package models

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

var format = "2006-01-02 15:04:05"

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey;comment:标识符"`
	CreatedAt JSONTime       `json:"created_at" gorm:"column:created_at;comment:创建于"`
	UpdatedAt JSONTime       `json:"updated_at" gorm:"column:updated_at;comment:更新于"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index; column:deleted_at;comment:删除于"`
}

type JSONTime struct {
	time.Time
}

func (baseModel *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	baseModel.CreatedAt = JSONTime{Time: time.Now()}
	baseModel.UpdatedAt = JSONTime{Time: time.Now()}
	return
}

func (baseModel *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	baseModel.UpdatedAt = JSONTime{Time: time.Now()}
	return
}

// MarshalJSON 重写序列化成json的时间格式
func (jsonTime JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", jsonTime.Format(format))
	return []byte(formatted), nil
}

func (jsonTime *JSONTime) UnmarshalJSON(data []byte) error {
	if len(data) < 1 {
		return nil
	}

	timeStr := string(data)
	timeStr = strings.Trim(timeStr, "\"")
	t, err := time.Parse(format, timeStr)
	*jsonTime = JSONTime{Time: t}

	return err

}

func (jsonTime JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if jsonTime.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return jsonTime.Time, nil
}

func (jsonTime *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*jsonTime = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
