# golang # gopkglist # http # gRPC # JWT
1.Get all gopkgs description and relationship between other pkgs data insert to mysql.  
2.Build a REST api using http server and gRPC service, JWT.  
.  
├── auth  
│   └── middleware.go -JWT auth  
├── conf  
│   └── app.ini -相關文件配置預設值  
├── controllers  
│   ├── gopkg.go -go package爬蟲  
│   └── user.go -使用者驗證  
├── debug  
│   └── http.go -開發用  
├── gopkg-spider -二進位執行檔  
├── helper  
│   └── utils.go   
├── main.go  
├── models  
│   ├── models.go -init 初始化db  
│   ├── pb_request.go -對protobuf service執行要求方法  
│   ├── query.go -sql指令  
│   ├── spider.go -爬蟲實作  
│   └── tables.go -定義資料表  
│   └── user.go -定義使用者  
├── pkg  
│   └── setting  
│       └── setting.go -調用配置文件  
├── protoserver  
│   ├── echospec -放置protobuf必要文件  
│   │   ├── route_pkgs.pb.go  
│   │   └── route_pkgs.proto  
│   └── server.go -protobuf server端實作  
└── routers  
    └── routers.go -http router 路由邏輯處理  
