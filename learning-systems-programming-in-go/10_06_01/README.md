# 10.6 FUSEを使った 作のファイルシステムの作成

1. FUSE インストール
    - https://osxfuse.github.io/
2. Apple Silicon Mac の場合はセキュリティ設定が必要（要再起動）
    - https://support.apple.com/ja-jp/guide/mac-help/mchl768f7291/mac
    
3. AWS S3 バケット作成
   1. AWSマネージメントコンソールのS3のコンソール画面でバケットを作成する
    - ```バケット名：go-system-study-202203```
    2. 何かファイルをアップロードしておく
4. AWS IAM Goup を作成
   - Group名: `Study`
   - アタッチするポリシー: AmazonS3ReadOnlyAccess
5. AWS IAM User作成
   1. `GoStudy` を作成する。
      1. 『Select AWS credential type』
         - [x] Access key - Programmatic access
         - [ ] Password - AWS Management Console access
    　
       1. 『Add user to group』 で `Study` グループを選択する 
   2. アクセスキーとシークレットアクセスキーはダウンロードするかメモしておく。
6.  ローカルマシンにAWSのキーを設定する
    1.  プロファイル名 `gostudy` でローカルマシンに設定
        ```shell
        cat << EOF > .envrc
        export AWS_PROFILE=gostudy
        EOF
        ```

        ```shell
        direnv allow .
        ```

        5 で生成したアクセスキーとシークレットキーを指定する
        ```shell
        $ aws configure
        AWS Access Key ID [****************XXXX]: xxxxx
        AWS Secret Access Key [****************XXXX]: xxxxx
        ```
7. マウントするローカルディレクトリ作成
```
mkdir -p /tmp/cloudfs
```
8. コンパイル実行

```shell
go build -o cloudfs main.go
./cloudfs s3://go-system-study-202203 /tmp/cloudfs
```

- 確認
```shell
$ ls -l /tmp/cloudfs
total 8
-r--r--r--  1 root  wheel  22  3 26 19:38 test.txt
$ file /tmp/cloudfs/test.txt
/tmp/cloudfs/test.txt: ASCII text
```

- ^C でcloudfsを終了させて確認するとファイルは見えなくなる
```shell
$ ls -l /tmp/cloudfs        
total 0
```

