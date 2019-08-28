package migratego

type QueryBuilder interface {
	DropTables(...string) DropTablesGenerator
	CreateTable(string, func(CreateTableGenerator)) CreateTableGenerator
	AlterTable(string, func(generator AlterTableGenerator)) AlterTableGenerator
	NewIndexColumn(column string, params ...interface{}) IndexColumnGenerator
	RawQuery(string)
	Sqls() []string
}

type Order string //todo  asc desc
type Querier interface {
	Sql() string
}
type DropTablesGenerator interface {
	Table(name string) DropTablesGenerator
	IfExists() DropTablesGenerator
	Sql() string
}
type CreateTableGenerator interface {
	Column(name string, Type string) TableColumnGenerator
	Index(name string, unique bool,indexType string) IndexGenerator
	Engine(engine string) CreateTableGenerator
	Charset(charset string) CreateTableGenerator
	Comment(comment string) CreateTableGenerator
	Sql() string
}
type TableColumnGenerator interface {
	GetName() string
	Primary(comment ...string) TableColumnGenerator
	NotNull() TableColumnGenerator
	Unsigned() TableColumnGenerator
	Binary() TableColumnGenerator
	ZeroFill() TableColumnGenerator
	Generated() TableColumnGenerator
	DefaultValue(v string) TableColumnGenerator
	Charset(charset string) TableColumnGenerator //todo charset
	Comment(c string) TableColumnGenerator
	AutoIncrement(primaryComment ...string) TableColumnGenerator
	Index(name string, unique bool,indexType string, params ...interface{}) IndexGenerator
	After(column string) TableColumnGenerator                                                 // todo after
	Rename(oldName string, newName string, charset string, collate string,notNull bool) TableColumnGenerator // todo after
	Sql() string
	GetPrimarySql() string
}
type AlterTableGenerator interface {
	Rename(name string) AlterTableGenerator
	Delete(name string) AlterTableGenerator
	AddColumn(name string, Type string) TableColumnGenerator
	RemoveColumn(name string) AlterTableGenerator
	AddIndex(name string, unique bool,indexType string,columns ... IndexColumnGenerator) IndexGenerator
	RemoveIndex(name string) AlterTableGenerator
	RemovePrimary(name string) AlterTableGenerator
	Charset(charset string) AlterTableGenerator // todo charset
	Comment(c string) AlterTableGenerator
	ModifyColumn(name string, Type string, notNull bool) TableColumnGenerator
	RenameColumn(oldName string, newName string, charset string,collate string, notNull bool) TableColumnGenerator
	Sql() string
}

type IndexGenerator interface {
	Unique() IndexGenerator
	Columns(...IndexColumnGenerator) IndexGenerator
	Sql() string
}
type IndexColumnGenerator interface {
	Sql() string
}

type DBClient interface {
	// PrepareTransactionsTable checks if table with migrations exists and creates it, if it doesn't
	PrepareTransactionsTable() error
	// Backup dumps database to some file in folder and returns path to it
	Backup(path string) (string, error)
	// InsertMigration adds migration to migrations table
	InsertMigration(migration *Migration) error
	// RemoveMigration removes migration from migrations table
	RemoveMigration(migration *Migration) error
	// ApplyMigration executes UpScript if down is false. Execute DownScript of down is true
	ApplyMigration(migration *Migration, down bool) error
	// GetAppliedMigrations returns list of migrations in migrations table
	GetAppliedMigrations() ([]Migration, error)
	//有错误是否继续
	SetContinueErr(flag bool)
}
