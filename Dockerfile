## http://www.iamlintao.com/6748.html
## final stage
#FROM alpine:3.8
##镜像创建作者
#MAINTAINER timel
##设置时区 【时区默认是美国时间】
##CentOS
#RUN echo "Asia/shanghai" > /etc/timezone
##Ubuntu
#RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
## 修改二进制加载问题
#RUN mkdir /lib64 && ln -s /lib/libc.musl-86_64.so.1 /lib64/ld-linux-x86-64.so.2
##创建文件
#RUN mkdir -p /server
##拷贝二进制docker
#COPY main /server/main
##跳转到指定目录
#WORKDIR /server
##设置程序权限
#RUN chmod +x main
##暴露端口
#EXPOSE 8081
##执行启动命令
#ENTRYPOINT ["./main"]

#**********************测试**************************#

#docker image build -t godemo .
#https://blog.csdn.net/hanyajun0123/article/details/90681253
# build stage
FROM golang:1.13.5 AS builder
#ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o godemo main.go

FROM alpine:3.8

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /go/release/godemo /
EXPOSE 3000
CMD ["/godemo"]



