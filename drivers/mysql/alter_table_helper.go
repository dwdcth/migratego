package mysql

import "github.com/dwdcth/migratego"

type RenameTableGenerator struct {
	oldName string
	newName string
}

func (s *RenameTableGenerator) Sql() string {
	return "RENAME TABLE " + wrapName(s.oldName) + " TO " + wrapName(s.newName)
}

const (
	AlterTableAdd  = "ADD"
	AlterTableDrop = "DROP"
)

type AlterTableGenHelper struct {
	table     string
	operation string
	query     migratego.Querier
}

type UpdateTableIndexes struct {
}

func (a *AlterTableGenHelper) Sql() string {
	sql := "ALTER TABLE " + wrapName(a.table)
	sql += " " + a.operation
	sql += " " + a.query.Sql()
	return sql
}
