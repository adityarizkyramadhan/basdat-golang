package model

import (
	"fmt"
	"time"
)

type Pelanggan struct {
	Id        int
	Name      string
	Address   string
	CreatedAt time.Time
}

func (*Pelanggan) TableName() string {
	return "pelanggan"
}

func (*Pelanggan) ColumnName() string {
	return "name, address, created_at"
}

func (p *Pelanggan) ExtractData() string {
	return fmt.Sprintf(`'%s', '%s', '%s'`, p.Name, p.Address, time.Now())
}
