package mysql

import (
	"strings"

	"github.com/dwdcth/migratego"
)

type IndexGenerator struct {
	name         string
	unique       bool
	columns      []migratego.IndexColumnGenerator
	parser       string
	keyBlockSize int
	indexType    string
	comment      string
}

func (i *IndexGenerator) Name(n string) migratego.IndexGenerator {
	i.name = n
	return i
}

func (i *IndexGenerator) Type(n string) migratego.IndexGenerator {
	i.name = n
	return i
}

func (i *IndexGenerator) Columns(c ...migratego.IndexColumnGenerator) migratego.IndexGenerator {
	i.columns = append(i.columns, c...)
	return i
}
func (i *IndexGenerator) Comment(c string) migratego.IndexGenerator {
	i.comment = c
	return i
}
func (i *IndexGenerator) Unique() migratego.IndexGenerator {
	i.unique = true
	return i
}
func (i *IndexGenerator) KeyBlockSize(s int) migratego.IndexGenerator {
	i.keyBlockSize = s
	return i
}
func (i *IndexGenerator) Parser(p string) migratego.IndexGenerator {
	i.parser = p
	return i
}
func (i *IndexGenerator) Sql() string {
	var sql string
	if len(i.columns) == 0 {
		return ""
	}
	if i.unique {
		sql += "UNIQUE "
	}
	columns := make([]string, len(i.columns))
	for i, c := range i.columns {
		columns[i] = c.Sql()
	}

	using := ""
	if i.indexType != "" {
		using += " USING " + i.indexType
	}
	sql += "INDEX " + wrapName(i.name) + using + "  ON (" + strings.Join(columns, ",") + ")"
	return sql
}

func newIndexGenerator(name string, unique bool, indexType string) migratego.IndexGenerator {
	result := &IndexGenerator{
		name:      name,
		unique:    unique,
		indexType: indexType,
	}
	return result
}
