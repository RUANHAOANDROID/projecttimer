# 使用 golang:latest 作为基础镜像
FROM golang:1.20-alpine

# 设置工作目录
WORKDIR /app

# 将当前目录下的所有文件复制到工作目录
COPY . .

# 使用 go mod 下载并安装依赖
RUN go mod download

# 构建 Go 应用程序
RUN go build -o main .

EXPOSE 6688

ENV APP_NAME pt

LABEL version="1.1"
# 暴露应用程序的端口
ENV APP_PORT 6688
#VOLUME ["/app/data"]
# 运行应用程序
CMD ["./main"]
#CMD ["./main", "-name", "$APP_NAME", "-port", "$APP_PORT"]
# --network bridge
# -p 6688:6688
# -v /mnt/disk1/appdata/pt/:/