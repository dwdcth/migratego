package mysql

import (
	"github.com/dwdcth/migratego"
	"strings"
)

type AlterTableGenerator struct {
	name    string
	columns []migratego.TableColumnGenerator
	indexes []migratego.IndexGenerator
	builder *MysqlQueryBuilder
}

func (t *AlterTableGenerator) AddColumn(name string, Type string) migratego.TableColumnGenerator {
	if name == "" {
		panic("Can't add column to table with empty name")
	}
	cg := &TableColumn{
		tableScope: t,
		name:       name,
		fType:      Type,
	}
	t.columns = append(t.columns, cg)
	return cg
}
func (t *AlterTableGenerator) RemoveColumn(name string) {
	q := rawQuery("COLUMN " + wrapName(name))
	g := AlterTableGenHelper{
		table:     t.name,
		operation: AlterTableDrop,
		query:     &q,
	}
	t.builder.generators = append(t.builder.generators, &g)
}

func (t *AlterTableGenerator) Rename(newName string) migratego.AlterTableGenerator {
	if newName == "" {
		panic("New name of table should not be empty")
	}
	t.builder.generators = append(t.builder.generators, &RenameTableGenerator{oldName: t.name, newName: newName})
	t.name = newName
	return t
}
func (t *AlterTableGenerator) AddIndex(name string, unique bool) migratego.IndexGenerator {
	index := newIndexGenerator(name, unique)

	g := AlterTableGenHelper{
		table:     t.name,
		operation: AlterTableAdd,
		query:     index,
	}
	t.builder.generators = append(t.builder.generators, &g)
	return index
}
func (t *AlterTableGenerator) RemoveIndex(name string) {
	q := rawQuery("INDEX " + wrapName(name))
	g := AlterTableGenHelper{
		table:     t.name,
		operation: AlterTableAdd,
		query:     &q,
	}
	t.builder.generators = append(t.builder.generators, &g)
}

//添加了comment
func (t *AlterTableGenerator) Comment(name string) migratego.AlterTableGenerator {
	q := rawQuery("COMMENT '" + name + "'")
	g := AlterTableGenHelper{
		table:     t.name,
		operation: "",
		query:     &q,
	}
	t.builder.generators = append(t.builder.generators, &g)
	return t
}

func (t *AlterTableGenerator) Charset(name string) migratego.AlterTableGenerator {
	q := rawQuery("COMMENT '" + name + "'")
	g := AlterTableGenHelper{
		table:     t.name,
		operation: "",
		query:     &q,
	}
	t.builder.generators = append(t.builder.generators, &g)
	return t
}

func (t *AlterTableGenerator) Delete(name string) {
	t.builder.generators = append(t.builder.generators, &dropTablesGenerator{tables: []string{t.name}})
}

func (t *AlterTableGenerator) ModifyColumn(name string, Type string, notNull bool) migratego.TableColumnGenerator {
	cg := &TableColumn{
		tableScope: t,
		name:       name,
		fType:      Type,
		notNull:    notNull,
		isModify:   true,
	}
	t.columns = append(t.columns, cg)
	return cg
}

func (t *AlterTableGenerator) RenameColumn(oldName string, newName string, charset string, collate string, notNull bool) migratego.TableColumnGenerator {
	cg := &TableColumn{isRename: true}
	cg.Rename(oldName, newName, charset, collate, notNull)
	t.columns = append(t.columns, cg)
	return cg
}

func (t *AlterTableGenerator) DeleteIfExists() {
	t.builder.generators = append(t.builder.generators, &dropTablesGenerator{ifExists: true, tables: []string{t.name}})
}

func (t *AlterTableGenerator) Sql() string {
	sql := "ALTER TABLE " + wrapName(t.name)
	sql += " " + strings.Join(t.builder.Sqls(), " ")

	//w := make([]string, len(c.columns))
	//for i, column := range c.columns {
	//	w[i] = column.Sql()
	//}
	return sql
}
