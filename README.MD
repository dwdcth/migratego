
forked from [github.com/saturn4er/migratego](https://github.com/saturn4er/migratego) ,origin [README](README_origin.MD)

## Examples

### Create table
```go
package main

import (
	"github.com/saturn4er/migratego"
	"github.com/saturn4er/migratego/types"
)


func init() {
	app.AddMigration(2, "initApp", initAppUp, initAppDown)
}
func initAppUp(s migratego.QueryBuilder) {
	s.CreateTable("user", func(t types.CreateTableGenerator) {
		t.Column("id", "int").Primary()
		t.Column("name", "varchar(255)").NotNull()
		t.Column("password", "varchar(255)").NotNull()

	})
}
func initAppDown(s migratego.QueryBuilder) {
	s.DropTables("user").IfExists()
}
```

## Sql generators
 * Drop tables
 * Create Table
   * Column
     * Primary
     * Not null
     * Unsigned
     * Binary
     * Zero Fill
     * Generated
     * Default Value
     * Comment
     * Index
     * Auto Increment
   * Index
   * Foreign key (TODO)
 * Table
   * Rename
   * Delete
   * Add column
   * Add index
   * Remove index
   * Update column
 * Raw Query
 

## TODO
- Add Postgres
- Foreign Key
- Tool to Create initial migration by existing tables
- Backup