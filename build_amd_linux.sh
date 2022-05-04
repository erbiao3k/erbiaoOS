CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" && upx erbiaoOS && scp erbiaoOS root@10.21.8.21:/ 
