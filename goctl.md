# 文档



## goctl代码自动生成命令

#### rpc 代码生成

```shell
goctl rpc proto -src add.proto -dir .
```

#### api代码生成

```shell
goctl api go -api bookstore.api -dir .
```

#### 定义数据库表结构，并生成CRUD+cache代码

```shell
goctl model mysql ddl -c -src book.sql -dir .
```





