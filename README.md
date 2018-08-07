# words-gin
## v0.2
### 需求分析
- 单词加入数据库
- 用户系统实现


### 版本改动：项目结构改变
```
├── cmd
│   └── myapp
│       └── main.go
├── service
│   ├── repository
│   │   ├── mysql
│   │   └── redis
│   └── word.go
├── transport
│   └── restful
│       └── word.go
└── word.go

```
