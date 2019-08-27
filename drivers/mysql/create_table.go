package mysql

import (
	"strings"

	"github.com/dwdcth/migratego"
)

type IndexType string

type Order string

type createTableGenerator struct {
	name          string
	engine        string
	comment       string
	charset       string
	columns       []migratego.TableColumnGenerator
	indexes       []migratego.IndexGenerator
	primaryKey    *PrimaryKeyGenerator
	uniqueIndexes map[string]string
}

func (c *createTableGenerator) Sql() string {
	var sql = "CREATE TABLE `" + c.name + "`("

	w := make([]string, len(c.columns))
	for i, column := range c.columns {
		w[i] = column.Sql()
	}
	sql += strings.Join(w, ",")

	if c.primaryKey != nil {
		sql += ", " + c.primaryKey.Sql()
	}
	for _, index := range c.indexes {
		indexSql := index.Sql()
		if indexSql != "" {
			sql += ", " + indexSql
		}
	}

	sql += ")"
	if c.engine != "" {
		sql += " ENGINE = " + c.engine
	}
	if c.charset != "" {
		sql += " DEFAULT CHARACTER SET = " + c.charset
	}
	if c.engine != "" {
		sql += " COMMENT = '" + strings.Replace(c.comment, "'", "\\'", -1) + "'"
	}
	return sql
}
func (c *createTableGenerator) Charset(charset string) migratego.CreateTableGenerator {
	c.charset = charset
	return c
}
func (c *createTableGenerator) Comment(comment string) migratego.CreateTableGenerator {
	c.comment = comment
	return c
}
func (c *createTableGenerator) Engine(engine string) migratego.CreateTableGenerator {
	c.engine = engine
	return c
}
func (c *createTableGenerator) Column(name string, Type string) migratego.TableColumnGenerator {
	result := &TableColumn{
		table: c,
		name:  name,
		fType: Type,
	}
	c.columns = append(c.columns, result)
	return result
}
func (c *createTableGenerator) Index(name string, unique bool) migratego.IndexGenerator {
	index := newIndexGenerator(name, unique)
	c.indexes = append(c.indexes, index)
	return index
}
func NewCreateTableGenerator(name string, sc func(migratego.CreateTableGenerator)) migratego.CreateTableGenerator {
	result := &createTableGenerator{
		name:          name,
		uniqueIndexes: make(map[string]string),
	}
	sc(result)
	return result
}
