package model

import (
	"fmt"
	"time"
)

type Bus struct {
	Id          int
	Name        string
	Quota       int
	Destination string
	Derpature   string
	CreatedAt   time.Time
}

func (*Bus) TableName() string {
	return "bus"
}

func (*Bus) ColumnName() string {
	return "name, quota, destination, derpature, created_at"
}

func (b *Bus) ExtractData() string {
	return fmt.Sprintf(`'%s',%d,'%s','%s','%s'`, b.Name, b.Quota, b.Destination, b.Derpature, time.Now())
}
