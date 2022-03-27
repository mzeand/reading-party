# 11.1.2 プロセスID

```shell

$ docker run -it --rm -v $(pwd):/go/src golang sh
# go version
go version go1.17.7 linux/arm64
# go build -o /go/bin/lspgo.11.1.2 /go/src/main.go
# # バックグラウンド実行
# /go/bin/lspgo.11.1.2 &
# プロセスID: 73
親プロセスID: 1

# # 親プロセス確認
# ps -p 1
  PID TTY          TIME CMD
    1 pts/0    00:00:00 sh
# # このプログラム実行のプロセス確認
# ps -p 73
  PID TTY          TIME CMD
   73 pts/0    00:00:00 lspgo.11.1.2
# exit

```
