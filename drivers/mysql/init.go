package mysql

import "github.com/dwdcth/migratego"

func init() {
	migratego.DefineDriver("mysql", QueryBuilderConstructor, NewClient)
}

func QueryBuilderConstructor() migratego.QueryBuilder {
	return &MysqlQueryBuilder{}
}
