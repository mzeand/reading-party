# 11.1.4 ユーザーIDとグループID

```shell

$ docker run -it --rm -v $(pwd):/go/src golang sh
# go version
go version go1.17.7 linux/arm64
# go build -o /go/bin/lsgo.11.1.4 /go/src/main.go
# whoami
root
# # 現在はroot
# lsgo.11.1.4
 ユーザーID: 0
 グループID: 0
 サブグループID: []
# # ユーザ追加
# useradd mizue
# # スイッチユーザ
# su mizue
$ lsgo.11.1.4
 ユーザーID: 1000
 グループID: 1000
 サブグループID: [1000]
$ exit
# # グループ追加してさっき追加したユーザをグループに追加
# groupadd study
# usermod -aG study mizue
# # スイッチユーザ（サブグループが追加されている）
# su mizue
$ lsgo.11.1.4
 ユーザーID: 1000
 グループID: 1000
 サブグループID: [1000 1001]
$ exit
# exit



```
