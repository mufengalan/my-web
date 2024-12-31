package model

type Example struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
}

// TableName 获取表名
func (m *Example) TableName() string {
	return "examples"
}
