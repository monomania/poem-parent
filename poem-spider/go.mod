module tesou.io/platform/poem-parent/poem-spider

go 1.13

replace (
	github.com/go-xorm/core v0.6.3 => github.com/go-xorm/core v0.6.2
	tesou.io/platform/poem-parent/poem-api => ../poem-api
	tesou.io/platform/poem-parent/poem-core => ../poem-core
)

require (
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/bitly/go-simplejson v0.5.0
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/hu17889/go_spider v0.0.0-20150809033053-85ede20bf88b
	golang.org/x/net v0.0.0-20191119073136-fc4aabc6c914
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
	tesou.io/platform/poem-parent/poem-api v1.0.0
	tesou.io/platform/poem-parent/poem-core v0.0.0-00010101000000-000000000000
)
