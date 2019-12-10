module tesou.io/platform/poem-parent/poem-core

go 1.13

replace (
	github.com/go-xorm/core v0.6.3 => github.com/go-xorm/core v0.6.2
	tesou.io/platform/poem-parent/poem-api => ../poem-api
)

require (
	github.com/astaxie/beego v1.12.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/core v0.6.3
	github.com/go-xorm/xorm v0.7.9
	github.com/kr/pretty v0.1.0 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	golang.org/x/text v0.3.1-0.20180807135948-17ff2d5776d2
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	tesou.io/platform/poem-parent/poem-api v1.0.0
)
