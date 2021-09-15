GOPROXY=https://goproxy.cn CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jarvis .
GOPROXY=https://goproxy.cn CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o jarvis-darwin .