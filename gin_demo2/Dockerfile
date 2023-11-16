FROM ubuntu:20.04

FROM golang:1.21.3
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

COPY .  /home/user/demo2 
RUN cd  /home/user/demo2 && go build .

EXPOSE 8081/tcp
WORKDIR /home/user/demo2
CMD ["./gin_demo2"]