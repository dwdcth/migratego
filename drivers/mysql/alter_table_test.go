package mysql

import (
	"testing"

	"github.com/dwdcth/migratego"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMysqlQueryBuilder_AlterTable(t *testing.T) {

	Convey("QueryBuilder.AlterTable Should generate right sql", t, func() {
		b := MysqlQueryBuilder{}
		b.AlterTable("test_table", func(g migratego.AlterTableGenerator) {
			id := g.AddColumn("id","int")
			So(id.GetName(), ShouldEqual, "id")
			t.Log(g.Sql())
			modify := g.ModifyColumn("id","varchar(10)",true).AutoIncrement()
			t.Log(modify.Sql())

			rename := g.RenameColumn("id","name","utf8mb4","utf8mb4_unicode_ci",true)
			t.Log(rename.Sql())
			//So(g.Sql(), ShouldEqual, "CREATE TABLE `test_table`(`id` varchar(255) NULL)")
			//id.Binary()
			//So(g.Sql(), ShouldEqual, "CREATE TABLE `test_table`(`id` varchar(255) BINARY NULL)")
			//id.Comment("test_comment")
			//So(g.Sql(), ShouldEqual, "CREATE TABLE `test_table`(`id` varchar(255) BINARY NULL COMMENT 'test_comment')")
			//id.DefaultValue("default_value")
			//So(g.Sql(), ShouldEqual, "CREATE TABLE `test_table`(`id` varchar(255) BINARY NULL DEFAULT 'default_value' COMMENT 'test_comment')")
			//id.Generated()
			//So(g.Sql(), ShouldEqual, "CREATE TABLE `test_table`(`id` varchar(255) BINARY NULL GENERATED ALWAYS AS ('default_value') COMMENT 'test_comment')")
			//id.NotNull()
			//So(g.Sql(), ShouldEqual, "CREATE TABLE `test_table`(`id` varchar(255) BINARY NOT NULL GENERATED ALWAYS AS ('default_value') COMMENT 'test_comment')")
			//id.Primary()
			//So(g.Sql(), ShouldEqual, "CREATE TABLE `test_table`(`id` varchar(255) BINARY NOT NULL GENERATED ALWAYS AS ('default_value') COMMENT 'test_comment', PRIMARY KEY (`id`))")
			//id.Primary("primary_comment")
			//So(g.Sql(), ShouldEqual, "CREATE TABLE `test_table`(`id` varchar(255) BINARY NOT NULL GENERATED ALWAYS AS ('default_value') COMMENT 'test_comment', PRIMARY KEY (`id`) COMMENT 'primary_comment')")
			//id.ZeroFill()
			//So(g.Sql(), ShouldEqual, "CREATE TABLE `test_table`(`id` varchar(255) ZEROFILL BINARY NOT NULL GENERATED ALWAYS AS ('default_value') COMMENT 'test_comment', PRIMARY KEY (`id`) COMMENT 'primary_comment')")
			//id.Unsigned()
			//So(g.Sql(), ShouldEqual, "CREATE TABLE `test_table`(`id` varchar(255) UNSIGNED ZEROFILL BINARY NOT NULL GENERATED ALWAYS AS ('default_value') COMMENT 'test_comment', PRIMARY KEY (`id`) COMMENT 'primary_comment')")
			//id.AutoIncrement()
			//So(g.Sql(), ShouldEqual, "CREATE TABLE `test_table`(`id` varchar(255) UNSIGNED ZEROFILL BINARY NOT NULL GENERATED ALWAYS AS ('default_value') AUTO_INCREMENT COMMENT 'test_comment', PRIMARY KEY (`id`))")
			//id.AutoIncrement("autoincrement_comment")
			//So(g.Sql(), ShouldEqual, "CREATE TABLE `test_table`(`id` varchar(255) UNSIGNED ZEROFILL BINARY NOT NULL GENERATED ALWAYS AS ('default_value') AUTO_INCREMENT COMMENT 'test_comment', PRIMARY KEY (`id`) COMMENT 'autoincrement_comment')")
			//id.Index("idx_id", true)
			//So(g.Sql(), ShouldEqual, "CREATE TABLE `test_table`(`id` varchar(255) UNSIGNED ZEROFILL BINARY NOT NULL GENERATED ALWAYS AS ('default_value') AUTO_INCREMENT COMMENT 'test_comment', PRIMARY KEY (`id`) COMMENT 'autoincrement_comment', UNIQUE INDEX `idx_id` (`id` ASC))")
		})
		g := b.DropTables("test_table").IfExists()
		So(g.Sql(), ShouldEqual, "DROP TABLE IF EXISTS `test_table`")
		g.Table("test_table_2")
		So(g.Sql(), ShouldEqual, "DROP TABLE IF EXISTS `test_table`,`test_table_2`")
		g = b.DropTables().IfExists()
		So(g.Sql(), ShouldEqual, "")
	})
}
