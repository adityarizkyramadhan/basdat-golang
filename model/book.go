package model

import (
	"fmt"
	"time"
)

type Book struct {
	Id        int
	IdUser    int
	IdBus     int
	CreatedAt time.Time
}

func (*Book) TableName() string {
	return "book"
}

func (*Book) ColumnName() string {
	return "id_user, id_bus, created_at"
}

func (b *Book) ExtractData() string {
	return fmt.Sprintf(`'%d','%d','%s'`, b.IdUser, b.IdBus, time.Now())
}
