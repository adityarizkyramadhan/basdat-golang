package querysql

import (
	"fmt"
)

func AddDataTable(data AddData) string {
	return fmt.Sprintf(`
	INSERT INTO %s (%s) VALUES (%s)
	`, data.TableName(), data.ColumnName(), data.ExtractData())
}

type AddData interface {
	TableName() string
	ColumnName() string
	ExtractData() string
}
