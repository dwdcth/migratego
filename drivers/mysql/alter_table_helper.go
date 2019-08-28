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
	AlterTableAdd    = "ADD"
	AlterTableDrop   = "DROP"
	AlterTableModify = "MODIFY COLUMN"
	AlterTableChange = "CHANGE"
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

	if colGen, ok := a.query.(migratego.TableColumnGenerator); ok && colGen.GetPrimarySql() != "" {
		sql += ", ADD " + colGen.GetPrimarySql()
	}
	return sql
}
