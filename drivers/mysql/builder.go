package mysql

import "github.com/dwdcth/migratego"

type MysqlQueryBuilder struct {
	generators []migratego.Querier
}

func (m *MysqlQueryBuilder) DropTables(names ...string) migratego.DropTablesGenerator {
	c := NewDropTablesGenerator(names...)
	m.generators = append(m.generators, c)
	return c
}
func (m *MysqlQueryBuilder) CreateTable(name string, g func(generator migratego.CreateTableGenerator)) migratego.CreateTableGenerator {
	c := NewCreateTableGenerator(name, g)
	m.generators = append(m.generators, c)
	return c
}
func (m *MysqlQueryBuilder) RawQuery(q string) {
	c := rawQuery(q)
	m.generators = append(m.generators, &c)
}

// NewIndexColumn creates new IndexColumnGenerator
// Usage NewIndexColumn(column, order[optional], length[optional])
// orderType default value is ASC
// length default value is int
func (c *MysqlQueryBuilder) NewIndexColumn(column string, params ...interface{}) migratego.IndexColumnGenerator {
	var length int
	var order = "ASC"
	var ok bool
	if len(params) > 0 {
		if order, ok = params[0].(string); !ok {
			panic("second param should be of type `string`")
		}
	}
	if len(params) > 1 {
		if length, ok = params[1].(int); !ok {
			panic("third param should be of type `int`")
		}
	}
	return &IndexColumnGenerator{
		Column: column,
		Order:  order,
		Length: length,
	}
}
func (m *MysqlQueryBuilder) AlterTable(name string, b func(t migratego.AlterTableGenerator)) migratego.AlterTableGenerator {

	c := NewAlterTableGenerator(name, b)
	m.generators = append(m.generators, c)
	return c
}
func (m *MysqlQueryBuilder) Sqls() []string {
	var result []string
	for _, g := range m.generators {
		result = append(result, g.Sql())
	}
	return result
}
