FROM golang:1.16 as build

WORKDIR /go/src/github.com/hippopos/jarvis/
COPY . .
RUN GOPROXY=https://goproxy.cn CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jarvis .

FROM alpine

RUN apk --no-cache add curl

ARG branch
ARG commit
ARG buildtime
ARG owner
ARG env_para
LABEL branch=$branch \
        commit=$commit \
        buildtime=$buildtime \
        maintainer=$owner \
        env_para=$env_para
LABEL email=cuiwenchang@k8sfans.com

RUN apk  add --no-cache tzdata
COPY --from=build /go/src/github.com/hippopos/jarvis/jarvis /
ENV TZ=Asia/Shanghai
EXPOSE 9999
EXPOSE 10000
WORKDIR /
RUN export GIN_MODE=release
ENTRYPOINT ["./jarvis"]
CMD ["server"]
