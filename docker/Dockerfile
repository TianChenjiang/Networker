# 源镜像
FROM golang:latest

# 开启go module
ENV GO111MODULE on
# 阿里云代理
ENV GOPROXY https://mirrors.aliyun.com/goproxy/
# ARGS
# 设置端口
ARG EXPOSED_PORT
# artificat名称, DockerFile放在根目录，arti名称作为子目录文件夹名称
ARG ARTIFACT_NAME


WORKDIR $GOPATH/src/github.com/$ARTIFACT_NAME/app

# 将服务器的go工程代码加入到docker容器中
ADD $ARTIFACT_NAME/. $GOPATH/src/github.com/$ARTIFACT_NAME/app

# 依赖装载 构建
RUN go build -o app
# 暴露端口
EXPOSE $EXPOSED_PORT
# 最终运行docker的命令
ENTRYPOINT ["./app"]