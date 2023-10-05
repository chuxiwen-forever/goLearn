## 初始化项目

- cd 03_bookstore
- go mod init github.com/liu/bookstore -> 初始化项目
- go get "github.com/jinzhu/gorm" -> 引入orm框架
- go get "github.com/jinzhu/gorm/dialects/mysql" -> mysql驱动
- go get "github.com/gorilla/mux" -> 构建web服务器