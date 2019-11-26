## 项目介绍
    项目尽可能的保留了自己在Java上的一些编码习惯，
    如三层结构：model层 service层，controller层. 
    希望可以对目前为Java开发者有需要转向go的朋友提供一个参考.

## 技术选型
* [xorm](https://github.com/go-xorm/xorm)
* [go_spider](https://github.com/hu17889/go_spider)
* [beego](https://github.com/astaxie/beego)

## 项目结构
~~~
poem-parent
|-- poem-api            实体类模块
|   |-- common          公共工具
|   |   `-- base        基础工具
|   `-- module          业务模块
|       `-- core
|-- poem-core           后台核心模块
|   |-- common          公共工具
|   |   |-- base        基础工具
|   |   |-- log         日志工具
|   |   `-- pinyin      拼音转换工具
|   |-- conf            配置文件
|   |-- module          业务模块            
|   |   `-- core
|   `-- test
|       `-- bson
|-- poem-spider         爬虫模块
|   |-- common          公共工具              
|   |   `-- base        基础工具
|   |-- conf            配置文件
|   |-- launch          爬虫启动类
|   `-- module          业务模块
|       `-- gushiwen
`-- poem-web            http服务模块
|    |-- common         公共工具
|    |   |-- base       基础工具
|    |   |-- fliters    过滤器        
|    |   `-- routers    路由配置
|    |-- conf           配置文件
|    |-- module         业务模块
|    |   |-- core
|    |   |-- index
|    |   `-- spider
|    `-- test
~~~
## 模块依赖
  
  | 模块名    |  依赖模块     |
  | --------    | :----:   |
  | poem-api  |无|
  | poem-core  |poem-api|
  | poem-spider  |poem-api oem-core|
  | poem-web  |poem-api poem-core poem-spider|
   


## 实现功能
* 诗人,诗文,诗句,古文,名句等数据的爬取
* http api 分页接口提供


    ```
    项目为一次尝试转换练习,主要是将java的编写习惯尝试转换到go的开发上
    ```

## 使用教程

* 配置环境
  * 安装 go
  * 环境变量
  
  | 变量名称=值    |  说明     |
  | --------    | :----:   |
  | GO111MODULE=on  |开启go mod模块支持|
  | GOPROXY=https://goproxy.cn,direct     |依赖包下载代理地址|
  | GOSUMDB=sum.golang.google.cn     |包的哈希值校验地址|
  
* 导入项目到[JetBrains GoLand](https://www.jetbrains.com/go/)
* 下载资源包
    ```
      cd ./poem-api
      go mod tidy
      cd ./poem-core
      go mod tidy
      cd ./poem-spider
      go mod tidy
      cd ./poem-web
      go mod tidy
    ```
* 手动创建数据库
  
  数据库名为: poem 
* 配置数据库连接
  
  各模块的conf下的 mysql.ini文件修改配置
* 同步数据库表

  运行入口: PC000Application.go
  ```
    注意运行时: working directory需为 ****/poem-parent/poem-core 下
  ```  
  ![](https://oscimg.oschina.net/oscnet/6aeea26d87faf8cc37c7a8de61d29f6c1e5.png "working directory")
* 执行爬取数据

   运行入口: PS000Application.go
   
* 启动http服务

   运行入口: PW000Application.go
  ![](https://oscimg.oschina.net/oscnet/b87398056bd5ffc0e7680f748c160bc7608.png "api")
  
* 联系作者
<table>
  <tr>
    <td>
    <td><img src="https://oscimg.oschina.net/oscnet/917bee8edddbf16a7645a56d085e887a59f.jpg"/></td> 
    <td><img src="https://oscimg.oschina.net/oscnet/aaf253aa4757b62af61036493f6fba683c2.jpg"/></td> 
    <td></td>
  </tr>
</table>
 