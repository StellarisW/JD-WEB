#获取golang
FROM golang:1.17 as go

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /www/wwwroot/www.stellaris.wang/JD-WEB

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件  可执行文件名为 app
RUN go build -o jd .

# 声明服务端口
EXPOSE 8080

# 启动容器时运行的命令
CMD ["./jd"]