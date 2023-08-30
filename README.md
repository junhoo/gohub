##goeasy-api

##### Build Setup
```
# 终端运行go程序
go run main.go

# 开启air 自动重载
air
```
##### Use Databse
```
# 创建database
mysql -u root -p

mysql> create database gohub; (Ps: uncreated use)
mysql> use gohub;
mysql> show tables;
```
---
##### Built-in
- gin
- viper（.env 和 config 信息的基础库）
- cast（内置安全的类型转换）
- gorm
- mysql
- sqlite
- go-redis
- zap（日志工具）
---