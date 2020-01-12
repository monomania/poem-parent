module tesou.io/platform/poem-parent/poem-web

replace (
	github.com/go-xorm/core v0.6.3 => github.com/go-xorm/core v0.6.2
	tesou.io/platform/poem-parent/poem-api => ../poem-api
	tesou.io/platform/poem-parent/poem-core => ../poem-core
	tesou.io/platform/poem-parent/poem-spider => ../poem-spider
)

go 1.13
