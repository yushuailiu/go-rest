## 项目
基于Gin的Go语言简单REST应用，Go入门学习示例、Go Web开发入门
## go安装
### Mac上安装
```shell
brew install go
```

### 环境变量配置
在~/.bashrc 或则 ~/.profile中添加
```shell
export GOROOT=/usr/local/Cellar/go/1.10/libexec  //具体参考自己安装的版本
export GOPATH=/Users/lfuture/Documents/goProject/go-rest:/Users/lfuture/Documents/goProject/其他项目地址 // 添加项目地址,第一个目录
填写你即将clone项目在你本地保存的目录
export GOBIN=/Users/lfuture/bin
```
使配置生效
```shell
source ~/.bashrc // 或则 source ~/.profile
```

## 配置及运行
```shell
git clone git@github.com:lfuture/go-rest.git
cd go-rest
go get github.com/gin-gonic/gin
```
### 创建数据库
```mysql
create database test;
CREATE TABLE `person` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(60) NOT NULL DEFAULT '',
  `last_name` varchar(60) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARSET=utf8
```
### 修改数据库
修改`mydatabase/mysql.go`
```go
SqlDB, err = sql.Open("mysql", "root:yourpassword@tcp(127.0.0.1:3306)/test?parseTime=true")
```
### 运行
```shell
go run main.go router.go //直接运行
go build                // 打包
```
### 访问
即可访问`router.go`里定义的相应路由～