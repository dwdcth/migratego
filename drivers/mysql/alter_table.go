package mysql

import (
	"github.com/dwdcth/migratego"
	"strings"
)

//https://blog.csdn.net/qq_36744627/article/details/81253482

type AlterTableGenerator struct {
	name       string
	indexes    []migratego.IndexGenerator
	builder    *MysqlQueryBuilder
	primaryKey *PrimaryKeyGenerator
}

func (t *AlterTableGenerator) AddColumn(name string, Type string) migratego.TableColumnGenerator {
	if name == "" {
		panic("Can't add column to table with empty name")
	}
	cg := &TableColumn{
		isAlter: true,
		name:    name,
		fType:   Type,
	}
	g := AlterTableGenHelper{
		table:     t.name,
		operation: AlterTableAdd,
		query:     cg,
	}

	t.builder.generators = append(t.builder.generators, &g)
	return cg
}
func (t *AlterTableGenerator) RemoveColumn(name string) migratego.AlterTableGenerator {
	q := rawQuery("COLUMN " + wrapName(name))
	g := AlterTableGenHelper{
		table:     t.name,
		operation: AlterTableDrop,
		query:     &q,
	}
	t.builder.generators = append(t.builder.generators, &g)
	return t
}

func (t *AlterTableGenerator) Rename(newName string) migratego.AlterTableGenerator {
	if newName == "" {
		panic("New name of table should not be empty")
	}
	t.builder.generators = append(t.builder.generators, &RenameTableGenerator{oldName: t.name, newName: newName})
	t.name = newName
	return t
}
func (t *AlterTableGenerator) AddIndex(name string, unique bool, indexType string, columns ...migratego.IndexColumnGenerator) migratego.IndexGenerator {
	index := newIndexGenerator(name, unique, indexType)
	index.Columns(columns...) // todo 类型判断
	g := AlterTableGenHelper{
		table:     t.name,
		operation: AlterTableAdd,
		query:     index,
	}
	t.builder.generators = append(t.builder.generators, &g)
	return index
}
func (t *AlterTableGenerator) RemoveIndex(name string) migratego.AlterTableGenerator {
	q := rawQuery("INDEX " + wrapName(name))
	g := AlterTableGenHelper{
		table:     t.name,
		operation: AlterTableDrop,
		query:     &q,
	}
	t.builder.generators = append(t.builder.generators, &g)
	return t
}
func (t *AlterTableGenerator) RemovePrimary(name string) migratego.AlterTableGenerator {
	q := rawQuery("PRIMARY " + wrapName(name))
	g := AlterTableGenHelper{
		table:     t.name,
		operation: AlterTableDrop,
		query:     &q,
	}
	t.builder.generators = append(t.builder.generators, &g)
	return t
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

func (t *AlterTableGenerator) Delete(name string) migratego.AlterTableGenerator {
	t.builder.generators = append(t.builder.generators, &dropTablesGenerator{tables: []string{t.name}})
	return t
}

func (t *AlterTableGenerator) ModifyColumn(name string, Type string, notNull bool) migratego.TableColumnGenerator {
	cg := &TableColumn{
		isAlter:  true,
		name:     name,
		fType:    Type,
		notNull:  notNull,
		isModify: true,
	}
	g := AlterTableGenHelper{
		table:     t.name,
		operation: AlterTableModify,
		query:     cg,
	}

	t.builder.generators = append(t.builder.generators, &g)
	return cg
}

func (t *AlterTableGenerator) RenameColumn(oldName string, newName string, charset string, collate string, notNull bool) migratego.TableColumnGenerator {
	cg := &TableColumn{isRename: true}
	cg.Rename(oldName, newName, charset, collate, notNull)

	g := AlterTableGenHelper{
		table:     t.name,
		operation: AlterTableChange,
		query:     cg,
	}
	t.builder.generators = append(t.builder.generators, &g)
	return cg
}

func (t *AlterTableGenerator) DeleteIfExists() {
	t.builder.generators = append(t.builder.generators, &dropTablesGenerator{ifExists: true, tables: []string{t.name}})
}

func (t *AlterTableGenerator) Sql() string {

	sql := strings.Join(t.builder.Sqls(), "; ")

	return sql
}

func NewAlterTableGenerator(name string, sc func(generator migratego.AlterTableGenerator)) migratego.AlterTableGenerator {
	result := &AlterTableGenerator{
		name:    name,
		builder: &MysqlQueryBuilder{},
	}
	sc(result)
	return result
}
