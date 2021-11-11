package main

import (
	"archive/zip"
	"os"
	"strings"
)

var source = "This archive contains some text files."

func main() {

	// https://pkg.go.dev/archive/zip

	// zipファイル作成
	zipFile, err := os.Create("test.zip")
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	f, err := zipWriter.Create("test.txt")
	if err != nil {
		panic(err)
	}

	// 文字列を読み込む
	reader := strings.NewReader(source)
	buf := make([]byte, reader.Size())
	_, err = reader.Read(buf)
	if err != nil {
		panic(err)
	}

	// zipに書き込む
	_, err = f.Write(buf)
	if err != nil {
		panic(err)
	}

}
