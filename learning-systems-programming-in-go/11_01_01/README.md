# 11.1.1 実行ファイル名

```shell

$ tree
.
└── main.go

$ docker run -it --rm -v $(pwd):/go/src golang sh
# go version
go version go1.17.7 linux/arm64
# ls -l /go/src
total 4
-rw-r--r-- 1 root root 192 Mar 27 01:49 main.go
# ls -l /go/bin
total 0
# go build -o /go/bin/lspgo-11.1.1 /go/src/main.go
# ls -l /go/bin
total 1808
-rwxr-xr-x 1 root root 1844838 Mar 27 04:14 lspgo-11.1.1
# cd /go/bin
# ./lspgo-11.1.1
 実行ファイル名: ./lspgo-11.1.1
 実行ファイルパス: /go/bin/lspgo-11.1.1
# exit


```
