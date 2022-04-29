# go-gin-gorm-oj

## git 环境

创建分支

```
git checkout -b fptang
```

更新分支代码并提交

```
git add *
git commit -m "first commit"
git push origin fptang
```



```
git init
git add README.md
git commit -m 
git branch -M main
git remote add origin git@github.com:tangfuping/go-oj.git
git push -u origin main
```



## go环境

在go-oj目录下：

```
go mod init go-oj
```



安装gin

GIN中文官网：[https://gin-gonic.com/zh-cn/docs/](https://gitee.com/link?target=https%3A%2F%2Fgin-gonic.com%2Fzh-cn%2Fdocs%2F)

```
go get -u github.com/gin-gonic/gin
```

安装gorm

GORM中文官网：[https://gorm.io/zh_CN/docs/](https://gitee.com/link?target=https%3A%2F%2Fgorm.io%2Fzh_CN%2Fdocs%2F)

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

```

Swagger

参考文档： [https://github.com/swaggo/gin-swagger](https://gitee.com/link?target=https%3A%2F%2Fgithub.com%2Fswaggo%2Fgin-swagger) 

接口访问地址：[http://localhost:8080/swagger/index.html](https://gitee.com/link?target=http%3A%2F%2Flocalhost%3A8080%2Fswagger%2Findex.html)

```go
go get -u github.com/swaggo/swag/cmd/swag
swag init
```

