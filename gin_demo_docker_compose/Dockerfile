FROM ubuntu:20.04

FROM golang:1.21.3
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

COPY .  /home/user/demo 
RUN cd  /home/user/demo && go build .

EXPOSE 8080/tcp
WORKDIR /home/user/demo 
CMD ["./gin_demo_docker_compose"]