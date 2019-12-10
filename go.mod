module tesou.io/platform/poem-parent

replace (
	github.com/go-xorm/core v0.6.3 => github.com/go-xorm/core v0.6.2
	tesou.io/platform/poem-parent/poem-api => ./poem-api
	tesou.io/platform/poem-parent/poem-core => ./poem-core
	tesou.io/platform/poem-parent/poem-spider => ./poem-spider
)

require (
	github.com/astaxie/beego v1.12.0
	github.com/lxn/walk v0.0.0-20191121152919-b7c43041fb1b // indirect
	github.com/lxn/win v0.0.0-20191106123917-121afc750dd3 // indirect
	gopkg.in/Knetic/govaluate.v3 v3.0.0 // indirect
	tesou.io/platform/poem-parent/poem-api v1.0.0
	tesou.io/platform/poem-parent/poem-core v0.0.0-00010101000000-000000000000
	tesou.io/platform/poem-parent/poem-spider v0.0.0-00010101000000-000000000000
)

go 1.13
