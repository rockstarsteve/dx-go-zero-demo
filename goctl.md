# 文档



## goctl代码自动生成命令

#### rpc 代码生成

```shell
goctl rpc proto -src .proto -dir .
```

#### api代码生成

```shell
goctl api go -api .api -dir .
```

#### 定义数据库表结构，并生成CRUD+cache代码

用sql脚本生成
```shell
goctl model mysql ddl -src .sql -dir . -c
```
用数据源生成
```shell
goctl model mysql datasource -url="数据库地址" -table="表明" -c -dir .
goctl model mysql datasource -url="root:root@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -table="sys_user" -c -dir . 
 
```





