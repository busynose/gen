package model

type Demo struct {
	Field1 string `gorm:"column:field_1"`
	Field2 int    `gorm:"column:field_2"`
	Field3 int64  `gorm:"column:field_3"`
}
