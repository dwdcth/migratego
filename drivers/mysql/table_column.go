package mysql

import (
	"strings"

	"github.com/dwdcth/migratego"
)

type CreateTableColumns []TableColumn

func (c CreateTableColumns) String() string {
	var result = make([]string, len(c))
	for i, v := range c {
		result[i] = "`" + strings.Replace(v.name, "`", "\\`", -1) + "`"
	}
	return strings.Join(result, ",")
}

type TableColumn struct {
	table      *createTableGenerator
	tableScope *AlterTableGenerator

	name     string
	oldName  string
	fType    string
	isModify bool
	isRename bool

	binary        bool
	unsigned      bool
	zeroFill      bool
	autoIncrement bool
	notNull       bool
	generated     bool
	comment       string
	defaultValue  string
	charset       string
	collate       string
	after         string
}

func (f *TableColumn) GetName() string {
	return f.name
}

// AutoIncrement define column as AUTO_INCREMENT and new PRIMARY INDEX
func (f *TableColumn) AutoIncrement(primaryComment ...string) migratego.TableColumnGenerator {
	f.autoIncrement = true
	if !strings.Contains(strings.ToLower(f.fType), "int") {
		panic(f.fType + " not support autoIncrement")
	}
	f.Primary(primaryComment...)
	return f
}

// Index add index to table for this column
// Usage: Index("index_name", true, "DESC", 10)
// This will create unique index "index_name" and it will add this column to it
func (f *TableColumn) Index(name string, unique bool, params ...interface{}) migratego.IndexGenerator {
	if name == "" {
		name = "idx_" + f.name
	}
	index := newIndexGenerator(name, unique)
	var order = "ASC"
	var length = 0
	if len(params) > 0 {
		o, ok := params[0].(string)
		if !ok {
			panic("Third param should be string (Order)")
		}
		if o != "" {
			order = o
		}
	}
	if len(params) > 1 {
		var ok bool
		length, ok = params[0].(int)
		if !ok {
			panic("Fourth param should be int (Length)")
		}

	}

	index.Columns(&IndexColumnGenerator{
		Column: f.name,
		Order:  order,
		Length: length,
	})
	f.table.indexes = append(f.table.indexes, index)
	return index
}
func (f *TableColumn) Primary(comment ...string) migratego.TableColumnGenerator {
	var c string
	if len(comment) > 0 {
		c = comment[0]
	}else {
		return f
	}
	f.table.primaryKey = NewPrimaryKeyGenerator([]string{f.name}, c)
	return f
}

// NotNull marks column as NOT NULL
func (f *TableColumn) NotNull() migratego.TableColumnGenerator {
	f.notNull = true
	return f
}

// Binary marks column as BINARY
func (f *TableColumn) Binary() migratego.TableColumnGenerator {
	f.binary = true
	return f
}
func (f *TableColumn) ZeroFill() migratego.TableColumnGenerator {
	f.zeroFill = true
	return f
}
func (f *TableColumn) Unsigned() migratego.TableColumnGenerator {
	f.unsigned = true
	return f
}
func (f *TableColumn) Generated() migratego.TableColumnGenerator {
	f.generated = true
	return f
}
func (f *TableColumn) DefaultValue(v string) migratego.TableColumnGenerator {
	f.defaultValue = v
	return f
}

func (f *TableColumn) Charset(v string) migratego.TableColumnGenerator {
	f.charset = v
	return f
}
func (f *TableColumn) Comment(v string) migratego.TableColumnGenerator {
	f.comment = v
	return f
}

func (f *TableColumn) After(filed string) migratego.TableColumnGenerator {
	f.after = filed
	return f
}

func (f *TableColumn) Rename(oldName string, newName string, charset string, collate string, notNull bool) migratego.TableColumnGenerator {
	f.oldName = oldName
	f.name = newName
	f.notNull = notNull
	f.charset = charset
	f.collate = collate
	return f
}

func (f *TableColumn) Sql() string {
	sql := "`" + f.name + "` " + string(f.fType)
	//sql := "ALTER TABLE " + wrapName(f.tableScope.name) + " ADD COLUMN " + wrapName(f.name) + " " + string(f.fType)

	if f.unsigned {
		sql += " UNSIGNED"
	}
	if f.zeroFill {
		sql += " ZEROFILL"
	}
	if f.binary {
		sql += " BINARY"
	}

	if f.defaultValue != "" {
		if f.generated {
			sql += " GENERATED ALWAYS AS ('" + strings.Replace(f.defaultValue, "'", "\\'", -1) + "')"
		} else {
			sql += " DEFAULT '" + strings.Replace(f.defaultValue, "'", "\\'", -1) + "'"
		}
	}
	if f.autoIncrement {
		sql += " AUTO_INCREMENT"
	}
	if f.charset != "" {
		sql += " CHARACTER SET '" + string(f.charset) + "'"
	}
	if f.collate != "" {
		sql += " COLLATE " + f.collate
	}
	if f.comment != "" {
		sql += " COMMENT '" + strings.Replace(f.comment, "'", "\\'", -1) + "'"
	}
	if f.after != "" {
		sql += " AFTER " + wrapName(f.after)
	}

	if f.notNull {
		sql += " NOT NULL"
	} else {
		sql += " NULL"
	}

	return sql
}
