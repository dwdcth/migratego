package main

import (
	"github.com/dwdcth/migratego"
	_ "github.com/dwdcth/migratego/drivers/mysql"
)

func init() {
	app.AddMigration(1, "initApp", initAppUp, initAppDown)
}
func initAppUp(s migratego.QueryBuilder) {
	s.CreateTable("user", func(t migratego.CreateTableGenerator) {
		t.Column("id", "int").Primary().Charset("aaa")
		t.Column("name", "varchar(255)").NotNull()
		t.Column("password", "varchar(255)").NotNull()
		t.Charset("utf8mb4")
	})
	s.AlterTable("user", func(t migratego.AlterTableGenerator) {
		t.RemoveColumn("1")
		t.AddColumn("aa","int").Charset("bbb")
		t.Sql()
	})
}
func initAppDown(s migratego.QueryBuilder) {
	s.DropTables("user").IfExists()
}
