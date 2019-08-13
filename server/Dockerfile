# 源镜像
FROM golang:latest
# 暂时一个app的gin web项目
WORKDIR $GOPATH/src/github.com/app

# 开启go module
ENV GO111MODULE on
# 阿里云代理
ENV GOPROXY https://mirrors.aliyun.com/goproxy/
# 将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/github.com/app
RUN export GO111MODULE=on
RUN go mod verify
RUN go build .
# 暴露端口
EXPOSE 8080
# 最终运行docker的命令
ENTRYPOINT ["./app"]
