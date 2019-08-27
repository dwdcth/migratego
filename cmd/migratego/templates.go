package main

const mainFileTpl = `package main

import (
	"os"

	"github.com/dwdcth/migratego"
	_ "github.com/dwdcth/migratego/drivers/{{.driver}}"
)

var app = migratego.NewApp("{{.driver}}", "{{.dsn}}")

func main() {
	{{ if .table }}app.SetSchemaVersionTable("{{.table}}") {{ end }}
	app.Run(os.Args)
}

`

const migrationFileTemplate = `package main

import (
	"github.com/dwdcth/migratego"
	{{ range $e := .imports }}
	"{{$e}}"
	{{ end }}
)

func init() {
	app.AddMigration(1, "{{.name}}", {{.name}}Up, {{.name}}Down)
}
func {{.name}}Up(s migratego.QueryBuilder) {
	{{.upBody}}
}
func {{.name}}Down(s migratego.QueryBuilder) {
	{{.downBody}}
}
`
