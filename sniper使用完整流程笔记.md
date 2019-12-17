#### GO安装及环境安装
GOPATH是GoLang的环境变量，GoRoot是go安装的地方，自动会生成，不要手动设置。GoBin是GoPath下的bin目录，主要用于放置一些第三方运行包。


#### 安装protoc
protoc是Protobuf 的编译器

brew tap grpc/grpc
brew install --with-plugins grpc

成功后使用
protoc --version
查看版本信息

可以查看路径/usr/local/bin/protoc


#### 安装protoc-gen-go
将proto文件编译成go代码时，需要用到protoc-gen-go插件
使用go get命令拉取代码及其依赖包，并自动完成编译和安装
// gRPC运行时接口编解码支持库
go get -u github.com/golang/protobuf/proto
// 从 Proto文件(gRPC接口描述文件) 生成 go文件 的编译器插件，执行完这句后可以在GoBin目录下看到protoc-gen-go可执行文件
go get -u github.com/golang/protobuf/protoc-gen-go


**使用**：将proto文件生成go文件，cd到存放当前.proto文件的目录下，执行：
protoc --go_out=plugins=grpc:. xxx.proto
**注意**：xxx是你要编译的.proto文件名，grpc:. xxx.proto冒号后面那个点和文件名之间有个空格，忘记加空格了不会编译通过，会报错：Missing input file.

#### 安装bilibili-twirp

//同样执行完后会在GoBin下有一个twirp的可执行文件
go get -u github.com/bilibili/twirp/protoc-gen-twirp

//twiter的源码地址
go get github.com/twitchtv/twirp/protoc-gen-twirp



整体使用: cd到存放当前.proto文件的目录下，执行
生成twirp.go和pb.go   命令: protoc --twirp_out=. --go_out=plugins=grpc:. test.proto

#### 代码书写：
.proto 定义接口和protocbuf数据

server.go 定义逻辑

http.go 注册接口

#### 代码解析：
.proto文件中：
service Echo定义的是接口。
每次写完proto需要重新编译成go文件









####  sniper使用

然后回到项目根目录，go build 编译一下，

//运行
go run main.go server --port=8080






处理代理：
git config --global http.proxy http://127.0.0.1:1087
git config --global https.proxy http://127.0.0.1:1087
