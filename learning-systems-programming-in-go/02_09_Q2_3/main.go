package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"name_en": "Hello World",
		"name_ja": "こんにちは 世界",
	}

	gzipw := gzip.NewWriter(w)
	defer gzipw.Close()

	encoder := json.NewEncoder(io.MultiWriter(gzipw, os.Stdout))
	encoder.SetIndent("", "     ")
	encoder.Encode(source)
	gzipw.Flush()

}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
