FROM golang
WORKDIR /go/src/app
COPY . .

# build steps
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN echo ">>> 1: go version" && go version \
&& echo ">>> 2: go mod tidy" && go mod tidy \
&& echo ">>> 3: go install" && go install

#原始方式：直接镜像内打包编译
RUN go build -o ./bin/ginsample
CMD ./bin/ginsample a