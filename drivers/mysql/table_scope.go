package mysql

import "github.com/saturn4er/migratego"

type TableScope struct {
	name    string
	builder *MysqlQueryBuilder
}
type UpdateTableIndexes struct {
}

func (t *TableScope) AddColumn(name string, Type string) migratego.UpdateTableAddColumnGenerator {
	if name == "" {
		panic("Can't add column to table with empty name")
	}
	cg := &UpdateTableAddColumnGenerator{
		tableScope: t,
		name:       name,
		fType:      Type,
	}
	t.builder.generators = append(t.builder.generators, cg)
	return cg
}
func (t *TableScope) RemoveColumn(name string) migratego.TableScope {
	q := rawQuery("COLUMN " + wrapName(name))
	g := AlterTableGenerator{
		table:     t.name,
		operation: AlterTableDrop,
		query:     &q,
	}
	t.builder.generators = append(t.builder.generators, &g)
	return t
}

func (t *TableScope) Rename(newName string) migratego.TableScope {
	if newName == "" {
		panic("New name of table should not be empty")
	}
	t.builder.generators = append(t.builder.generators, &RenameTableGenerator{oldName: t.name, newName: newName})
	t.name = newName
	return t
}
func (t *TableScope) AddIndex(name string, unique bool) migratego.IndexGenerator {
	index := newIndexGenerator(name, unique)

	g := AlterTableGenerator{
		table:     t.name,
		operation: AlterTableAdd,
		query:     index,
	}
	t.builder.generators = append(t.builder.generators, &g)
	return index
}
func (t *TableScope) RemoveIndex(name string) migratego.TableScope {
	q := rawQuery("INDEX " + wrapName(name))
	g := AlterTableGenerator{
		table:     t.name,
		operation: AlterTableAdd,
		query:     &q,
	}
	t.builder.generators = append(t.builder.generators, &g)
	return t
}
//添加了comment
func (t *TableScope) Comment(name string) migratego.TableScope {
	q := rawQuery("COMMENT '" + name + "'")
	g := AlterTableGenerator{
		table:     t.name,
		operation: "",
		query:     &q,
	}
	t.builder.generators = append(t.builder.generators, &g)
	return t
}
func (t *TableScope) Delete() {
	t.builder.generators = append(t.builder.generators, &dropTablesGenerator{tables: []string{t.name}})
}
func (t *TableScope) DeleteIfExists() {
	t.builder.generators = append(t.builder.generators, &dropTablesGenerator{ifExists: true, tables: []string{t.name}})
}
