 protoc -I../a -I. --go_out=paths=source_relative:.  *.proto

 # 注意 此处生成的b.pb.go 需要变更引用a.pb.go的路径 目前我并没有找到简单的解决办法 生产环境直接用脚本修改掉 import的包
