protoc -I ./include/ -I ./ --go_out=./ *.proto
# 注意：如果在protoc中使用了-I选项 一定要加上-I 选项 说明自己proto文件的目录
# 因为proto他不明白
